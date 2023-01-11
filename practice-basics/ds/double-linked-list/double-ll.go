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

func (d *doublyLinkedList) prepend(value int) {
	newNode := node{data: value}
	if d.head == nil {
		d.head = &newNode
		d.length++
		return
	}
	d.head.prev = &newNode
	newNode.next = d.head
	d.head = &newNode
	d.length++
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
		fmt.Printf("prev = %v value=%v next =%v addr =%v\n", curr.prev, curr.data, curr.next, &curr)
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

func main() {
	myList := doublyLinkedList{}
	myList.prepend(10)
	myList.prepend(20)
	myList.prepend(30)
	myList.prepend(40)
	fmt.Println("List:")
	myList.print()
	myList.deleteWithValue(30)
	myList.deleteWithValue(40)
	fmt.Println("List:")
	myList.print()
	myList.prepend(30)
	myList.prepend(40)
	myList.deleteWithValue(50)
	fmt.Println("List:")
	myList.print()
	myList.postpend(50)
	fmt.Println("List:")
	myList.print()
}
