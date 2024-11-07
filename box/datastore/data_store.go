package datastore

import (
	"sync"
	"yushu/box/config"
	"yushu/box/utility/singleton"
)

const (
	DSDefault uint8 = iota
	DSMutex
	DSRWMutex
)

type DsInterface interface {
	Get(key string) (value interface{}, err error)
}

type DataStore struct {
	_config   *config.Config
	_key      string
	_value    interface{}
	_dataZone *DataHashMap
	_queue    chan struct{}
	_mutex    *sync.Mutex
	_rwMutex  *sync.RWMutex
}

var dataLazySingleton singleton.Lazy

func New() *DataStore {
	ins := dataLazySingleton.Instance(&DataStore{})
	return (*ins).(*DataStore)
}

func (ds *DataStore) Config() *config.Config {
	return ds._config
}

// Put 存入队列
func (ds *DataStore) Put() {
	ds._queue <- struct{}{}
}

// Leave 释放队列
func (ds *DataStore) Leave() {
	<-ds._queue
}

func Get(key string) (interface{}, bool) {
	return New()._dataZone.Get(key)
}

func Set(key string, value interface{}) {
	New()._dataZone.Add(key, value)
}

// init 初始化 datastore
func init() {
	conf := config.New()
	if conf == nil {
		panic("config is nil")
	}
	ds := New()
	ds._config = conf
	ds._queue = make(chan struct{}, conf.Queue.MaxConnNum)
	ds._dataZone = NewHashMap()
	// 配置文件
	ds._dataZone.Add("config", conf)
}
