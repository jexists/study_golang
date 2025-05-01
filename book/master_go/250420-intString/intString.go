package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 100
	input := strconv.Itoa(n)
	fmt.Println(input)
	input = strconv.FormatInt(int64(n), 10)
	fmt.Println(input)
	input = string(n)
	fmt.Println(input)
	// 100
	// 100
	// d
	n = 100
	input = strconv.Itoa(n)
	fmt.Println(input)
	input = strconv.FormatInt(int64(n), 10)
	fmt.Println(input)
	input = string(n)
	fmt.Println(input)
	// 10
	// 10
	//

}
