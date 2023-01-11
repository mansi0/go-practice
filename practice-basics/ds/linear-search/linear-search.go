package main

import "fmt"

func linearS(arr []int, key int) (int, bool) {
	for i, num := range arr {
		if num == key {
			return i, true
		}
	}
	return 0, false
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}

	fmt.Println(linearS(arr1, 3))
	fmt.Println(linearS(arr1, 10))

	arr2 := []int{-1, -2, -3, 4, 7}
	fmt.Println(linearS(arr2, -3))
	fmt.Println(linearS(arr2, 3))

}
