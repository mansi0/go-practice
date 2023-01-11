package main

import "fmt"

func findSum(arr []int, sum int) {

	for i:= 0 ; i< len(arr)-1 ; i++ {
		for j:= i+1 ; j< len(arr) ; j++ {
			if arr[i] + arr[j] == sum {
				fmt.Println("ans:",arr[i],arr[j])
				return
			}
		}
	}
	fmt.Println("ans:",-1)
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
	var sum int
	fmt.Println("Enter Sum")
	fmt.Scanf("%d", &sum)
	findSum(arr, sum)
}
