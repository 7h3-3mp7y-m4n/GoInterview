package hashmap

import "fmt"

type KeyValue struct {
	key   int
	value string
	next  *KeyValue
}

type hashMap struct {
	buckets []*KeyValue
	size    int
}

func main() {
	hashmap := NewHashMap(5)
	hashmap.Put(1, "apple")
	hashmap.Put(8, "carrot")
	hashmap.Put(11, "banana")
	val, ok := hashmap.Get(8)
	if ok {
		fmt.Println("Key -> 6 ", val)
	}

	//typical dymanic leetcode hashmap
	m := make(map[string]int)
	m["apple"] = 2
	m["kiwi"] = 1
	val1, ok := m["apple"]
	if ok {
		fmt.Println("apple count -- ", val1)
	}
	delete(m, "apple")
	for k, v := range m {
		fmt.Println(k, v)
	}

}

func NewHashMap(size int) *hashMap {
	return &hashMap{
		buckets: make([]*KeyValue, size),
		size:    size,
	}
}

func (h *hashMap) hashFunction(key int) int {
	return key % h.size
}

// Put inserts or updates a key-value pair
func (h *hashMap) Put(key int, value string) {
	index := h.hashFunction(key)
	head := h.buckets[index]

	// check if key already exists -> update
	for node := head; node != nil; node = node.next {
		if node.key == key {
			node.value = value
			return
		}
	}

	// otherwise insert new node at head
	newNode := &KeyValue{key: key, value: value, next: head}
	h.buckets[index] = newNode
}

func (h *hashMap) Get(key int) (string, bool) {
	index := h.hashFunction(key)
	for node := h.buckets[index]; node != nil; node = node.next {
		if node.key == key {
			return node.value, true
		}
	}
	return "", false
}
