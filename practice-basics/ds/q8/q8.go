package main

import "fmt"

func findOccurance(str string) {
	h1 := make(map[string]int)
	for _, i := range str {
		h1[string(i)]++
	}
	for key, cnt := range h1 {
		fmt.Printf("%v%d ", key, cnt)
	}
}

func main() {
	str := "elephant"
	findOccurance(str)
}
