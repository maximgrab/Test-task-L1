package main

import "fmt"

func deleteSliceElement(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}
func main() {
	slice := []string{"cat", "cat", "dog", "cat", "tree"}
	slice = deleteSliceElement(slice, 3)
	fmt.Println(slice)
}
