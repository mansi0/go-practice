package main

import "fmt"

type stack []string

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) push(str string) {
	*s = append(*s, str)
}

func (s *stack) pop() (string, bool) {
	if s.isEmpty() {
		return "", false
	}
	element := (*s)[len(*s)-1]

	*s = (*s)[:len(*s)-1]
	return element, true
}

func main() {
	var s1 stack
	s1.push("Hello")
	s1.push("Data Structure")
	fmt.Println(s1.pop())
}
