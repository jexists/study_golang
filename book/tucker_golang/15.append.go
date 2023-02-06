package main

import "fmt"

func main() {
	var slice1 = []int{1, 2, 3}
	// 한개 추가
	slice2 := append(slice1, 4)
	fmt.Println(slice1)
	// [1 2 3]
	fmt.Println(slice2)
	// [1 2 3 4]

	// 여러개 추가
	slice3 := append(slice1, 4, 5, 6, 7, 8)
	fmt.Println(slice3)
	// [1 2 3 4 5 6 7 8]

	var slice = []int{1, 2, 3}
	for i := 1; i < 5; i++ {
		slice = append(slice, i)
	}
	fmt.Println(slice)
	// [1 2 3 1 2 3 4]
	slice = append(slice, 6, 7, 8)
	fmt.Println(slice)
	// [1 2 3 1 2 3 4 6 7 8]
}
