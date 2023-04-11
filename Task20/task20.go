package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var res strings.Builder
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = str[:len(str)-1]
	splited := strings.Split(str, " ")
	for i := 0; i < len(splited)-1; i++ {
		res.WriteString(splited[len(splited)-i-1] + " ")
	}
	res.WriteString(splited[0])
	fmt.Println(res.String())
}
