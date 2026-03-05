package main

import (
	"fmt"
)

func sumArray() {
	var num int

	fmt.Print("N=")

	fmt.Scan(&num)

	arr := make([]int64, num)

	for i := 0; i < num; i++ {
		fmt.Scan(&arr[i])
	}
	sum := 0
	for i := 0; i < num; i++ {
		sum += int(arr[i])
	}
	fmt.Println(sum)
}

// Slice
func main() { // Find maxiumum element from array
	// var n int
	// fmt.Scan(&n)
	// arr := make([]int64, n)

	// for i := 0; i < n; i++ {
	// 	fmt.Scan(&arr[i])
	// }
	// max := arr[0]

	// for i := 0; i < len(arr); i++ {
	// 	if max < arr[i] {
	// 		max = arr[i]
	// 	}
	// }
	//fmt.Println(max)
	// mySlice := make([]int, 2, 5)
	// mySlice1 := make([]int, 5)
	// fmt.Println(mySlice)
	// fmt.Println(mySlice1)

	// nums := []int{1, 2, 3}
	// nums = append(nums, 4, 5)
	// fmt.Println(nums)
	sumArray()
}
