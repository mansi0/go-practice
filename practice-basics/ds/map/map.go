package main

import "fmt"

func main() {
	var m map[string]string
	m = map[string]string{"hi": "hello"}
	fmt.Println(m)
	delete(m, "hi")
	fmt.Println(m)
}
