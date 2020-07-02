package main

import "fmt"

func printArray(arr []int) {
	//数组的遍历
	arr[0] = 100

	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	var arr2 = [3]int{1, 3, 5}
	arr3 := [...]int{4, 5, 2, 3, 1}

	var grid = [3][2]int{}
	//var grid [3][2]int{}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(grid)

	// 数组传值
	printArray(arr1[:])
	fmt.Println()
	//printArray(arr2)
	printArray(arr3[:])

	fmt.Println(arr1, arr3)

}
