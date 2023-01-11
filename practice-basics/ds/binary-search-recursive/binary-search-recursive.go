package main

import "fmt"

func binaryS_recc(arr []int, key int, low int, high int) (int, bool) {
	if low <= high {
		mid := (low + high) / 2
		fmt.Println(mid)
		if arr[mid] == key {
			fmt.Println("inside if")
			return mid, true
		} else if key < arr[mid] {
			return binaryS_recc(arr, key, low, mid-1)
		} else {
			return binaryS_recc(arr, key, mid+1, high)
		}
	}
	return 0, false

}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	key, ans := binaryS_recc(arr1, 9, 0, len(arr1)-1)
	fmt.Println(key, ans)
	if ans {
		fmt.Println("key is found at position", key)
	} else {
		fmt.Println("Key not found")
	}
}
