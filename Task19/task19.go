package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	var res strings.Builder
	fmt.Scan(&str)
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		res.WriteString(string(runes[len(runes)-i-1]))
	}
	fmt.Println(res.String())
}
