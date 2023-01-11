package main

import (
	"fmt"
)

func join(separator string, st ...string) string {
	var stringjoin string
	for _, s := range st[:len(st)-1] {
		stringjoin += string(s) + separator
	}
	return stringjoin + st[len(st)-1]
}

func main() {
	s1 := join("@", "Hello", "World", "Welcome!")
	fmt.Println(s1)

	s2 := join("", "")
	fmt.Println(s2)

}
