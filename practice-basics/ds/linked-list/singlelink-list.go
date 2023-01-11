package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkdlist struct {
	head   *node
	length int
}

func (l *linkdlist) prepend(value int) {
	newNode := node{data: value}
	if l.head != nil {
		newNode.next = l.head
		l.head = &newNode
		l.length++
	} else {
		l.head = &newNode
		l.length++
	}
}

func (l *linkdlist) postpend(value int) {
	newNode := node{data: value}
	if l.length == 0 {
		l.head = &newNode
		l.length++
		return
	}
	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = &newNode
	l.length++
}

func (l *linkdlist) printList() {
	if l.head == nil {
		return
	}
	currNode := l.head
	for currNode != nil {
		fmt.Println(currNode.data)
		currNode = currNode.next
	}
}

func (l *linkdlist) deleteWithValue(value int) {
	if l.head == nil {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}
	next := l.head.next
	prev := l.head
	for next != nil {
		if next.data == value {
			prev.next = next.next
			l.length--
			return
		}
		prev = next
		next = next.next
	}
}

func main() {
	myList := linkdlist{}
	myList.prepend(10)
	myList.prepend(20)
	myList.prepend(30)
	myList.prepend(40)
	myList.printList()
	myList.deleteWithValue(30)
	myList.deleteWithValue(40)
	myList.printList()
	myList.prepend(30)
	myList.prepend(40)
	myList.deleteWithValue(50)
	myList.printList()
	myList.postpend(50)
	myList.printList()
}
