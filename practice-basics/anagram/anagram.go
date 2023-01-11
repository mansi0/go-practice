package main

import "fmt"

func checkIfStringAnagram(s string, d string) bool {
	if len(s) != len(d) {
		return false
	}
	source := make(map[string]int)
	destination := make(map[string]int)

	for _, l := range s {
		source[string(l)]++
	}
	for _, l := range d {
		destination[string(l)]++
	}

	for l, cnt := range source {
		if destination[l] != cnt {
			return false
		}
	}
	return true
}

func main() {
	var s, t string
	fmt.Println("Enter Source string:")
	fmt.Scanln(&s)
	fmt.Println("Enter Target String:")
	fmt.Scanln(&t)
	fmt.Println(checkIfStringAnagram(s, t))
}
