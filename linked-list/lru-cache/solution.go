package lrucache

type node struct {
	key, value int
	prev, next *node
}

type LRUCache struct {
	cap   int
	cache map[int]*node
	left  *node // The least recently used.
	right *node // The most recently used.
}

// Remove the node from the list.
func (lru *LRUCache) remove(node *node) {
	prev, next := node.prev, node.next
	prev.next = next
	next.prev = prev
}

// Insert the node at the right most position.
func (lru *LRUCache) insert(node *node) {
	prev, next := lru.right.prev, lru.right
	prev.next, next.prev = node, node
	node.next = next
	node.prev = prev
}

func Constructor(capacity int) LRUCache {
	left, right := &node{}, &node{}
	left.next = right
	right.prev = left

	return LRUCache{
		cap:   capacity,
		cache: make(map[int]*node, capacity),
		left:  left,
		right: right,
	}
}

func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.cache[key]; ok {
		lru.remove(node) // Remove current node from the LRU cache.
		lru.insert(node) // Re-insert current node at the right most position.

		return node.value
	}

	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	if node, ok := lru.cache[key]; ok {
		lru.remove(node)
	}

	lru.cache[key] = &node{key: key, value: value}
	lru.insert(lru.cache[key])

	if len(lru.cache) > lru.cap {
		// Evict the least recently used node.
		evicted := lru.left.next
		lru.remove(evicted)
		delete(lru.cache, evicted.key)
	}
}
