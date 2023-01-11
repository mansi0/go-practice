package main

import "fmt"

func binary_search(arr []int, key int) (int, bool) {
	low := 0
	high := len(arr) - 1
	for low <= high {

		mid := (low + high) / 2
		if arr[mid] == key {
			return mid, true
		} else if key < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0, false
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	key, ans := binary_search(arr1, 5)
	if ans {
		fmt.Println("key is found at position", key)
	}
}
