package app

import "yushu/box/datastore"

func Register(key string, value interface{}) {
	datastore.Set(key, value)
}

func init() {
	
}
