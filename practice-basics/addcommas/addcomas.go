package main

import (
	"bytes"
	"fmt"
)

func concatenateI(input string) string {

	var buffer bytes.Buffer

	for index, i := range input {
		if string(i) == " " && (index != len(input)-1) {
			buffer.WriteString(", ")
		} else {
			buffer.WriteRune(i)
		}
	}
	return buffer.String()
}

func main() {
	fmt.Println(concatenateI("1 2 3"))
	fmt.Println(concatenateI(""))
}
