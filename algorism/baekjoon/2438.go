package main

import (
	"fmt"
	"strings"
)

func main() {

	var num int
	fmt.Scanf("%d", &num)

	for i := 1; i <= num; i++ {
		str := "*"
		fmt.Println(strings.Repeat(str, i))
	}
}
