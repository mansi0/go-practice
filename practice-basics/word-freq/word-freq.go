package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// dat, err := os.ReadFile("word-freq/data.txt")
	// check(err)
	// fmt.Print(string(dat))

	f, err := os.Open("word-freq/data.txt")
	check(err)

	// b1 := make([]byte, 5)
	// n1, err := f.Read(b1)
	// check(err)
	// fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	reader := bufio.NewScanner(bufio.NewReader(f))
	reader.Split(bufio.ScanWords)

	wordcnt := make(map[string]int)

	for reader.Scan() {
		wordcnt[reader.Text()]++
	}

	for index, k := range wordcnt {
		fmt.Println(index, k)
	}
}
