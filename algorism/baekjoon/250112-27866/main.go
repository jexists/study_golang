package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	var n int
	fmt.Scan(&n)

	fmt.Println(string([]rune(s)[n-1]))
}
