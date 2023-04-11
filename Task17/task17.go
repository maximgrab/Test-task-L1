package main

import "fmt"

func binarySearch(array []int, key int) (bool, int) {
	left := 0
	right := len(array) - 1
	for left <= right {
		mid := left + (right-left)/2
		if array[mid] == key {
			return true, mid
		}
		if array[mid] < key {
			left = mid + 1
		}
		if array[mid] > key {
			right = mid - 1
		}
	}
	return false, 0
}
func main() {
	array := []int{0, 2, 4, 5, 6, 7, 9, 12, 30}
	fmt.Println(binarySearch(array, 12))
	fmt.Println(binarySearch(array, 30))
	fmt.Println(binarySearch(array, 0))
	fmt.Println(binarySearch(array, 235))
	fmt.Println(binarySearch(array, -2))
}
