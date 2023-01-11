package main

import "fmt"

func getInvCount(arr []int) int {
	cnt := 0
	for i:= 0 ; i< len(arr) ; i++ {
		for j:= i+1 ; j< len(arr) ; j++ {
			if arr[i] > arr[j] {
				cnt += 1
			}
		}
	}
	return cnt
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
	fmt.Println("Inversion count is:", getInvCount(arr))
}
