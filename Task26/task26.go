package main

import (
	"fmt"
	"strings"
)

func main() {
	uniqeMap := make(map[rune]struct{})
	str := ""
	fmt.Scan(&str)
	str = strings.ToLower(str)
	for _, i := range str {
		if _, ok := uniqeMap[i]; ok {
			fmt.Println(false)
			return
		}
		uniqeMap[i] = struct{}{}
	}
	fmt.Println(true)
}
