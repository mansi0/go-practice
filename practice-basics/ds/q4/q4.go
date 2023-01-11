package main

import "fmt"

type node struct {
	data int
	prev *node
	next *node
}

type doublyLinkedList struct {
	head   *node
	length int
}

func (d *doublyLinkedList) postpend(value int) {
	newNode := node{data: value}
	if d.head == nil {
		d.head = &newNode
		d.length++
		return
	}
	curr := d.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = &newNode
	newNode.prev = curr
	d.length++
}

func (d *doublyLinkedList) print() {
	if d.length == 0 {
		return
	}
	curr := d.head
	for curr != nil {
		fmt.Printf("%v<->", curr.data)
		curr = curr.next
	}
}

func (d *doublyLinkedList) deleteWithValue(value int) {
	if d.length == 0 {
		return
	}
	if d.head.data == value {
		d.head = d.head.next
		d.head.prev.next = nil
		d.head.prev = nil
		d.length--
		return
	}
	curr := d.head
	for curr != nil {
		if curr.data == value {
			curr.prev.next = curr.next
			curr.next.prev = curr.prev
			d.length--
			return
		}
		curr = curr.next
	}
}

func (d *doublyLinkedList) search(key int) string {
	curr := d.head
	for curr != nil {
		if curr.data == key {
			return "YES"
		}
		curr = curr.next
	}
	return "NO"
}

func main() {
	myList := doublyLinkedList{}
	myList.postpend(1)
	myList.postpend(2)
	myList.postpend(3)
	myList.postpend(4)
	fmt.Println("List:")
	myList.print()

	fmt.Println(myList.search(4))
	fmt.Println(myList.search(5))
}
