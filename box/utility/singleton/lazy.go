package singleton

import "sync"

// Lazy 单例结构
type Lazy struct {
	ins  *interface{}
	mu   sync.Mutex
	init bool
}

// Instance 单例实例
func (singleton *Lazy) Instance(inter interface{}) *interface{} {
	if !singleton.init {
		singleton.mu.Lock()
		defer singleton.mu.Unlock()
		if !singleton.init {
			singleton.ins = &inter
			singleton.init = true
		}
	}
	return singleton.ins
}
