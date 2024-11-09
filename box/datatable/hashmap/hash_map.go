package hashmap

import (
	"errors"
	"fmt"
	"sync"
	"yushu/box/logger"
)

// node 节点结构，每个节点代表一个键值对，使用链表解决哈希冲突
type node struct {
	key    string
	value  interface{}
	next   *node
	wMutex sync.RWMutex
}

// HashMap 哈希表结构
type HashMap struct {
	buckets []*node      // 桶数组，存放链表头节点
	size    int          // 当前桶数量
	count   int          // 当前键值对数量
	mutex   sync.RWMutex // 全局互斥锁用于扩容
}

// 设置装载因子的阈值
const loadFactorThreshold = 0.75

// NewHashMap 创建一个新的哈希表
func NewHashMap(size ...int) *HashMap {
	// 默认初始容量为16
	sizeNum := 16
	if len(size) > 0 {
		sizeNum = size[0]
	}
	return &HashMap{
		buckets: make([]*node, sizeNum),
		size:    sizeNum,
	}
}

// hash 计算键的哈希值
func (h *HashMap) hash(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = (31*hash + int(key[i])) % h.size
	}
	return hash
}

// resize 扩容并重新散列所有键
func (h *HashMap) resize() {
	newSize := h.size * 2
	newBuckets := make([]*node, newSize)

	for _, head := range h.buckets {
		for curr := head; curr != nil; curr = curr.next {
			newIndex := h.hash(curr.key) % newSize
			newNode := &node{key: curr.key, value: curr.value, next: newBuckets[newIndex]}
			newBuckets[newIndex] = newNode
		}
	}
	h.buckets = newBuckets
	h.size = newSize
}

// Set 向哈希表中插入键值对
func (h *HashMap) Set(key string, value interface{}) {
	// value 不能为 nil
	if value == nil {
		logger.Info(logger.ErrorType, fmt.Sprintf("hash map key = %v, value can not be nil", key))
		return
	}
	h.mutex.Lock()
	defer h.mutex.Unlock()

	// 检查是否需要扩容
	if float64(h.count)/float64(h.size) > loadFactorThreshold {
		h.resize()
	}

	index := h.hash(key)
	head := h.buckets[index]

	// 如果键已经存在于链表中，更新它的值
	for curr := head; curr != nil; curr = curr.next {
		if curr.key == key {
			// 使用读写锁保护对值的修改
			curr.wMutex.Lock()
			curr.value = value
			// 释放写锁
			curr.wMutex.Unlock()
			return
		}
	}

	// 否则将新节点插入链表头部
	newNode := &node{key: key, value: value, next: head}
	h.buckets[index] = newNode
	h.count++
}

// Get 从哈希表中查找键的值
func (h *HashMap) Get(key string) (interface{}, error) {
	index := h.hash(key)
	head := h.buckets[index]

	for curr := head; curr != nil; curr = curr.next {
		if curr.key == key {
			return curr.value, nil
		}
	}
	return nil, errors.New(key + ": not found")
}

// Delete 从哈希表中删除键
func (h *HashMap) Delete(key string) bool {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	index := h.hash(key)
	head := h.buckets[index]

	if head == nil {
		return false
	}

	// 如果头节点就是要删除的键
	if head.key == key {
		h.buckets[index] = head.next
		h.count--
		return true
	}

	// 遍历链表找到要删除的节点
	prev := head
	for curr := head.next; curr != nil; curr = curr.next {
		if curr.key == key {
			prev.next = curr.next
			h.count--
			return true
		}
		prev = curr
	}
	return false
}
