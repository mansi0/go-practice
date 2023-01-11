package main

import "fmt"

type stack []string

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) push(str string) {
	*s = append(*s, str)
}

func (s *stack) pop() string {
	if s.isEmpty() {
		return "STack is Empty"
	}
	element := (*s)[len(*s)-1]

	*s = (*s)[:len(*s)-1]
	return element
}

func main() {
	var s1 stack
	str1 := "reverse me"
	for _, i := range str1 {
		s1.push(string(i))
	}
	for !s1.isEmpty() {
		fmt.Printf(s1.pop())
	}
}
