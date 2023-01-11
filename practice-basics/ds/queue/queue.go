package main

import "fmt"

type queue []string

func (q *queue) isEmpty() bool {
	return len(*q) == 0
}

func (q *queue) enQueue(str string) {
	*q = append(*q, str)
}

func (q *queue) deQueue() (string, bool) {
	if q.isEmpty() {
		return "", false
	}
	element := (*q)[0]

	*q = (*q)[1:]
	return element, true
}

func main() {
	var q1 queue
	q1.enQueue("Hello")
	q1.enQueue("Data Structure")
	for !q1.isEmpty() {
		fmt.Println(q1.deQueue())
	}
}
