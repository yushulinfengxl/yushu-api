package datastore

import (
	"fmt"
	"sync"
	"yushu/box/utility/singleton"
)

// DataEntry 表示哈希表中的单个条目，通过双向链表连接到其他条目。
type DataEntry struct {
	Key   string      // 键
	Value interface{} // 值
	_last *DataEntry  // 指向前一个条目的指针
	_next *DataEntry  // 指向后一个条目的指针
}

var dataEntryLazySingleton singleton.Lazy

// NewDataEntry 创建并返回一个新的 DataEntry 实例。
func NewDataEntry() *DataEntry {
	ins := dataEntryLazySingleton.Instance(&DataEntry{})
	return (*ins).(*DataEntry)
}

// DataHashMap 表示整个哈希表结构。
type DataHashMap struct {
	head *DataEntry // 链表的头节点
	mu   sync.Mutex // 互斥锁，保证并发安全
}

var dataHashMapEntryLazySingleton singleton.Lazy

// NewHashMap 创建并返回一个新的 DataHashMap 实例。
func NewHashMap() *DataHashMap {
	ins := dataHashMapEntryLazySingleton.Instance(&DataHashMap{})
	return (*ins).(*DataHashMap)
}

// Add 将一个新的键值对插入到哈希表中。
func (hm *DataHashMap) Add(key string, value interface{}) {
	hm.mu.Lock()         // 加锁以确保线程安全
	defer hm.mu.Unlock() // 函数结束时解锁

	newEntry := &DataEntry{
		Key:   key,
		Value: value,
	}

	if hm.head == nil {
		// 如果链表为空，将新条目设为头节点
		hm.head = newEntry
	} else {
		// 将新条目插入到链表末尾
		current := hm.head
		for current._next != nil {
			if current.Key == key {
				// 如果键已存在，则更新其值
				current.Value = value
				return
			}
			current = current._next
		}
		if current.Key == key {
			// 如果链表末尾的键已存在，则更新其值
			current.Value = value
		} else {
			// 否则，将新条目添加到链表末尾
			current._next = newEntry
			newEntry._last = current
		}
	}
}

// Get 根据键从哈希表中检索值。
func (hm *DataHashMap) Get(key string) (interface{}, bool) {
	hm.mu.Lock()         // 加锁以确保线程安全
	defer hm.mu.Unlock() // 函数结束时解锁

	current := hm.head
	for current != nil {
		if current.Key == key {
			return current.Value, true // 如果找到键，返回值和 true
		}
		current = current._next
	}
	return nil, false // 如果未找到键，返回 nil 和 false
}

// Remove 从哈希表中删除一个键值对。
func (hm *DataHashMap) Remove(key string) {
	hm.mu.Lock()         // 加锁以确保线程安全
	defer hm.mu.Unlock() // 函数结束时解锁

	current := hm.head
	for current != nil {
		if current.Key == key {
			if current._last != nil {
				// 将前一个节点的 _next 指向当前节点的 _next
				current._last._next = current._next
			} else {
				// 如果当前节点是头节点，则将头节点指向下一个节点
				hm.head = current._next
			}
			if current._next != nil {
				// 将下一个节点的 _last 指向当前节点的 _last
				current._next._last = current._last
			}
			return
		}
		current = current._next
	}
}

// PrintAll 打印哈希表中的所有键值对（用于调试）。
func (hm *DataHashMap) PrintAll() {
	current := hm.head
	for current != nil {
		fmt.Printf("Key: %s, Value: %v\n", current.Key, current.Value)
		current = current._next
	}
}
