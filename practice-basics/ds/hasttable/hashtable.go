package main

import "fmt"

const arraysize = 7

type hashTable struct {
	array [arraysize]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key  string
	next *bucketNode
}

func initHashTable() *hashTable {
	result := &hashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func (h *hashTable) insert(key string) {
	index := hash(key)
	h.array[index].insertBucket(key)
}

func (h *hashTable) search(key string) bool {
	index := hash(key)
	return h.array[index].searchBucket(key)
}
func (h *hashTable) delete(key string) {
	index := hash(key)
	h.array[index].deleteBucketNode(key)
}

func (h *hashTable) print() {
	for i := range h.array {
		h.array[i].printBucket()
		fmt.Println()
	}

}

func hash(key string) int {
	sum := 0
	for _, k := range key {
		sum += int(k)
	}
	return sum % arraysize
}

func (b *bucket) insertBucket(k string) {
	if !b.searchBucket(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println("Already Exist")
	}
}

func (b *bucket) searchBucket(k string) bool {
	curr := b.head
	for curr != nil {
		if curr.key == k {
			return true
		}
		curr = curr.next
	}
	return false
}

func (b *bucket) deleteBucketNode(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	curr := b.head
	for curr.next != nil {
		if curr.next.key == k {
			curr.next = curr.next.next
			return
		}
		curr = curr.next
	}

}

func (b *bucket) printBucket() {
	curr := b.head
	for curr != nil {
		fmt.Print(curr.key, "\t")
		curr = curr.next
	}
}

func main() {
	testHashTable := initHashTable()
	list := []string{
		"Eric",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}
	for _, v := range list {
		testHashTable.insert(v)
	}
	testHashTable.delete("BUTTERS")
	testHashTable.print()
	testHashTable.delete("STAN")
	testHashTable.print()
	fmt.Println(testHashTable.search("BUTTERS"))
}
