package main
import "fmt"


func rearrange(arr []int) {
	for i := 1;i<len(arr);i++ {
		if(arr[i-1]<0 && arr[i]<0) {
			k := searchForPositive(arr,i+1)
			if(k!=-1) {
				swap(&arr,i,k)
			}
		} else if(arr[i-1]> 0 && arr[i] > 0){
			k := searchForNegative(arr,i+1)
			if(k!=1) {
				swap(&arr,i,k)
			}
		}
	}
}

func searchForPositive(arr []int,i int)(int) {
	for ; i<len(arr);i++ {
		if(arr[i]>0) {
			return i
		}
	}
	return -1
}


func searchForNegative(arr []int,i int)(int) {
	for ;i<len(arr);i++ {
		if(arr[i]<0) {
			return i
		} 
	}
	return 1
}

func swap(arr *[]int,i,k int) {
	temp := (*arr)[i]
	(*arr)[i] = (*arr)[k]
	(*arr)[k] = temp
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
	rearrange(arr)
	fmt.Println(arr)
}