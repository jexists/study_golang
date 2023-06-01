package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	for i, v := range a {
		fmt.Printf("index: %d / value: %d \n", i, v)
		// index: 0 / value: 1
		// index: 1 / value: 2
		// index: 2 / value: 3
		// index: 3 / value: 4
		// index: 4 / value: 5
	}
}
