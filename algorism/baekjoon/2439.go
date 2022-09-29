package main

import (
	"fmt"
	"strings"
)

func main() {
	var num int
	fmt.Scanf("%d", &num)

	for i := 1; i <= num; i++ {
		space := strings.Repeat(" ", num-(i))
		string := strings.Repeat("*", i)
		fmt.Println(space + string)
	}
}
