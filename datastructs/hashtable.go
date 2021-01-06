package datastructs

import "fmt"

// ArraySize is the size of the Hashtable array
const ArraySize = 7

// HashTable struct
type HashTable struct {
	array [ArraySize]*bucket
}

// bucket struct
type bucket struct {
	head *bucketNode
}

// bucketNode struct
type bucketNode struct {
	key  string
	next *bucketNode
}

// Insert, takes a key and adds it to the table
func (h *HashTable) Insert(key string) {
	idx := hash(key)
	h.array[idx].insert(key)
}

// Search take in a key and return whether it is in the table or not
func (h *HashTable) Search(key string) bool {
	idx := hash(key)
	return h.array[idx].search(key)
}

// Delete will return true if the item has been deleted.
func (h *HashTable) Delete(key string) {
	idx := hash(key)
	h.array[idx].delete(key)
	return
}

// insert takes a key and adds it to a bucket
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "already exists")
	}
}

// search searches for a node in the bucket
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// delete deletes an item in the bucket
func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			previousNode.next = previousNode.next.next
			return
		}
		previousNode = previousNode.next
	}
}

// hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// InitHashTable
func InitHashTable() *HashTable {
	res := &HashTable{}
	for i := range res.array {
		res.array[i] = &bucket{}
	}
	return res
}
