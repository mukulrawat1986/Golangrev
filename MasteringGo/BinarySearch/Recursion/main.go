// Binary search implementation in Go, using recursion

package main

import "fmt"

func binarySearch(array []int, target int, lowIndex int, highIndex int) int {

	if highIndex < lowIndex {
		return -1
	}

	mid := int((lowIndex + highIndex) / 2)

	if array[mid] > target {
		return binarySearch(array, target, lowIndex, mid-1)
	} else if array[mid] < target {
		return binarySearch(array, target, mid+1, highIndex)
	} else {
		return mid
	}
}

func main() {

	arr := []int{3, 4, 12, 13, 35, 46, 78, 90, 89, 111, 127}

	res := binarySearch(arr, 4, 0, 10)
	fmt.Printf("%d\n", res)

	res = binarySearch(arr, 89, 0, 10)
	fmt.Printf("%d\n", res)

	res = binarySearch(arr, 150, 0, 10)
	fmt.Printf("%d\n", res)
}
