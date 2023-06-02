package main

import "fmt"

type myInt int

func (a myInt) add(b int) int {
	// return a + b // ERROR
	return int(a) + b
}

func main2() {
	var a myInt = 10
	fmt.Println(a.add(30))
	var b int = 20
	// fmt.Println(b.add(50))  // ERROR
	fmt.Println(myInt(b).add(50))
}
