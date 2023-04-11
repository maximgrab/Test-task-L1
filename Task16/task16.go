package main

import "fmt"

func quickSort(array []int, low, hight int) {
	if low >= hight || low < 0 {
		return
	}
	p := partition(array, low, hight)
	quickSort(array, low, p-1)
	quickSort(array, p+1, hight)

}
func partition(array []int, low, hight int) int {
	pivot := array[hight]
	i := low - 1
	for j := low; j < hight; j++ {
		if array[j] <= pivot {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}
	i++
	array[i], array[hight] = array[hight], array[i]
	return i
}
func main() {
	array := []int{0, 6, 4, 2, 9, 5, 6, 7, 4, 2, 1}
	quickSort(array, 0, len(array)-1)
	fmt.Println(array)
}
