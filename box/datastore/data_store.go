package datastore

import (
	"sync"
	"yushu/box/datatable/hashmap"
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
	_key      string
	_value    interface{}
	_dataZone *hashmap.HashMap
	_queue    chan struct{}
	_mutex    *sync.Mutex
	_rwMutex  *sync.RWMutex
}

var dataLazySingleton singleton.Lazy

func New() *DataStore {
	ins := dataLazySingleton.Instance(&DataStore{
		_dataZone: hashmap.NewHashMap(),
	})
	return (*ins).(*DataStore)
}

func Set(key string, value interface{}) {
	ds := New()
	ds._dataZone.Set(key, value)
}

func Get(key string) (value interface{}, err error) {
	ds := New()
	return ds._dataZone.Get(key)
}

// Put 存入队列
func (ds *DataStore) Put() {
	ds._queue <- struct{}{}
}

// Leave 释放队列
func (ds *DataStore) Leave() {
	<-ds._queue
}
