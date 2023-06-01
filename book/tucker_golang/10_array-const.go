package main

import "fmt"

const Y int = 3

func main() {
	// x := 5
	const z = 7
	// a := [x]int{1, 2, 3, 4, 5} ERROR
	b := [Y]int{1, 2, 3} // [1 2 3]
	c := [z]int{1, 2, 3} // [1 2 3 0 0 0 0]
	var d [10]int        // [0 0 0 0 0 0 0 0 0 0]
	// fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
