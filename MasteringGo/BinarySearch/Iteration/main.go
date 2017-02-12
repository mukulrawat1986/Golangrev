// Binary search impleemntation in Go using iteration

package main

import "fmt"

func iterBinarySearch(array []int, target int, lowIndex int, highIndex int) int {
	startIndex := lowIndex
	endIndex := highIndex
	var mid int
	for startIndex <= endIndex {
		mid = int((startIndex + endIndex) / 2)
		if array[mid] > target {
			endIndex = mid - 1
		} else if array[mid] < target {
			startIndex = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	arr := []int{3, 4, 12, 13, 35, 46, 78, 90, 89, 111, 127}

	res := iterBinarySearch(arr, 4, 0, 10)
	fmt.Printf("%d\n", res)

	res = iterBinarySearch(arr, 89, 0, 10)
	fmt.Printf("%d\n", res)

	res = iterBinarySearch(arr, 150, 0, 10)
	fmt.Printf("%d\n", res)
}
