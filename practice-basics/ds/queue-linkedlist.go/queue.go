package main

import "fmt"

type node struct {
	data int
	next *node
}

type queueList struct {
	head   *node
	length int
}

func (q *queueList) isEmpty() bool {
	return q.length == 0
}

func (q *queueList) enQueue(value int) {
	newNode := node{data: value}
	if q.head == nil {
		q.head = &newNode
		q.length++
		return
	}
	curr := q.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = &newNode
	q.length++
}

func (q *queueList) deQueue() (int, bool) {
	if q.isEmpty() {
		return 0, false
	}
	ele := q.head.data
	q.head = q.head.next
	q.length--
	return ele, true
}

func main() {
	q1 := queueList{}
	q1.enQueue(1)
	q1.enQueue(2)
	q1.enQueue(3)
	for !q1.isEmpty() {
		fmt.Println(q1.deQueue())
	}
}
