package main

import "fmt"

type LinkedList struct {
	key   int
	value int
	prev  *LinkedList
	next  *LinkedList
}

type LRUCache struct {
	totalCapacity   int
	currentCapacity int
	mapping         map[int]*LinkedList
	head            *LinkedList
	tail            *LinkedList
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		totalCapacity: capacity,
		mapping:       map[int]*LinkedList{},
	}
	return cache
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.mapping[key]
	if !ok {
		return -1
	}
	this.priorityNode(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.mapping[key]
	if !ok {
		node = &LinkedList{key: key, value: value}
		this.addNode(node)
		if this.currentCapacity > this.totalCapacity {
			this.deleteHead()
		}
	} else {
		node.value = value
		this.priorityNode(node)
	}
}

func (this *LRUCache) addNode(node *LinkedList) {
	this.mapping[node.key] = node
	this.currentCapacity += 1
	if this.head == nil && this.tail == nil {
		this.head = node
		this.tail = node
		return
	}
	this.tail.next = node
	node.prev = this.tail
	node.next = nil
	this.tail = node
}

func (this *LRUCache) deleteHead() {
	delete(this.mapping, this.head.key)
	this.currentCapacity -= 1
	if this.currentCapacity == 0 {
		this.head = nil
		this.tail = nil
	} else {
		this.head.next.prev = nil
		this.head = this.head.next
	}
}

func (this *LRUCache) priorityNode(node *LinkedList) {
	if this.tail == node {
		return
	}
	node.next.prev = node.prev
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		this.head = this.head.next
	}
	this.tail.next = node
	node.prev = this.tail
	node.next = nil
	this.tail = node
}

func (this *LRUCache) debug() {
	keys := []int{}
	values := []int{}
	cur := this.head
	for cur != nil {
		keys = append(keys, cur.key)
		values = append(values, cur.value)
		cur = cur.next
	}
	fmt.Println("[DEBUG]", values)

	keys = []int{}
	values = []int{}
	cur = this.tail
	for cur != nil {
		keys = append(keys, cur.key)
		values = append(values, cur.value)
		cur = cur.prev
	}
	fmt.Println("[DEBUG]", values)
}

func main() {
	obj := Constructor(3)
	obj.Put(1, 1)
	obj.Put(2, 2)
	obj.Put(3, 3)
	obj.Put(4, 4)
	fmt.Println(obj.Get(4))
	obj.debug()
	fmt.Println(obj.Get(3))
	obj.debug()
	fmt.Println(obj.Get(2))
	obj.debug()
	fmt.Println(obj.Get(1))
	obj.Put(5, 5)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(2))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))
	fmt.Println(obj.Get(5))
}
