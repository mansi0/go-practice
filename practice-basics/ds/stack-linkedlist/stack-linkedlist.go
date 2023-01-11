package main

import "fmt"

type node struct {
	data int
	next *node
}

type stackList struct {
	head   *node
	length int
}

func (s *stackList) isEmpty() bool {
	return s.length == 0
}

func (s *stackList) push(value int) {
	s.head = &node{value, s.head}
	s.length++
}

func (s *stackList) pop() (int, bool) {
	if s.isEmpty() {
		return 0, false
	}
	ele := s.head.data
	s.head = s.head.next
	s.length--
	return ele, true
}

func main() {
	s1 := stackList{}
	s1.push(1)
	s1.push(2)
	s1.push(3)
	for !s1.isEmpty() {
		fmt.Println(s1.pop())
	}
}
