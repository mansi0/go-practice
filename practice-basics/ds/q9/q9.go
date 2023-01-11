package main

import "fmt"

func findCLosest(arr []int,n int,target int)(int){
	if(target <= arr[0]) {
		return arr[0]
	}
	if(target >= arr[len(arr)-1]) {
		return arr[n-1]
	}

	i := 0
	j := n
	var mid int
	for i<j {
		mid = (i+j) /2
		if(arr[mid] == target) {
			return arr[mid]
		}
		if(target <arr[mid]) {
			if(mid > 0 && target>arr[mid-1]) {
				return getClosest(arr[mid-1],arr[mid],target)
			}
			j = mid
		} else {
			if(mid<len(arr)-1 && target < arr[mid+1]) {
				return getClosest(arr[mid],arr[mid+1],target)
			}
			i = mid+1
		}
	}
	return arr[mid]
}

func getClosest(v1,v2,target int)int {
	if(target-v1 >= v2-target) {
		return v2
	} else {
		return v1
	}
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
	var t int
	fmt.Println("Enter target number")
	fmt.Scanf("%d", &t)
	fmt.Println(arr)
	fmt.Println("Output:", findCLosest(arr,n,t))
}
