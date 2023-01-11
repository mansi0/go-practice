package main

import "fmt"

func bubbleSort(arr *[]int, n int) int {

	for i := 1; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				swap := (*arr)[i]
				(*arr)[i] = (*arr)[j]
				(*arr)[j] = swap
			}
		}
		rounds = i
		if isSorted(*arr, n) {
			return i
		}
	}
	return rounds
}

func isSorted(arr []int, n int) bool {
	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Println("Enter size of array")

	fmt.Scanf("%d", &n)
	arr := make([]int, n)
	fmt.Println("Enter array elements:")
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Println("Rounds :", bubbleSort(&arr, n))
}
