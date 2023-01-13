package main

import "fmt"

//reverse slice in the same slice

func main() {
	var arr = []int{1, 2, 3, 4, 5}

	//for i := range arr {
	//	defer fmt.Println(arr[i])
	//}
	//h:=int(len(arr)/2)-1
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	fmt.Println(arr)
}
