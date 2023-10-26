package main

import (
	"fmt"
)

// func main() {
// 	x := [...]int{10, 20, 30}
// 	y := [...]int{}

// 	fmt.Println(x, cap(x), len(x))
// 	// [10 20 30] 3 3
// 	fmt.Println(y, cap(y), len(y))
// 	// [] 0 0

// 	x = [4]int{1, 2, 3, 4}
// 	x = [...]int{1, 2, 3, 4}
// }

// ================================
// type Student1 struct {
// 	A int8
// 	B int
// 	C int8
// 	D int
// 	E int8
// }

// type Student2 struct {
// 	B int
// 	A int8
// 	D int
// 	C int8
// 	E int8
// }

// type Student3 struct {
// 	B int
// 	D int
// 	A int8
// 	C int8
// 	E int8
// }

// type Student4 struct {
// 	A int8
// 	C int8
// 	E int8
// 	B int
// 	D int
// }

// func main() {

// 	var student1 Student1
// 	var student2 Student2
// 	var student3 Student3
// 	var student4 Student4

// 	fmt.Println(unsafe.Sizeof(student1))
// 	fmt.Println(unsafe.Sizeof(student2))
// 	fmt.Println(unsafe.Sizeof(student3))
// 	fmt.Println(unsafe.Sizeof(student4))

// }

func main() {
	var x = new(int)
	fmt.Println(x == nil)
	// false
	fmt.Println(*x)
	// 0
	var y = new(string)
	fmt.Println(y == nil)
	// false
	fmt.Println(*y == "")
	// true
}
