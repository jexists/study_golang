package main

import "fmt"

func main() {
	var array [10]int
	fmt.Println(array)
	// [0 0 0 0 0 0 0 0 0 0]

	var slice []int
	if len(slice) == 0 {
		fmt.Println("slice empty ", slice)
		// slice empty  []
	}
	// slice[1] = 10 // ERROR
	// panic: runtime error: index out of range [1] with length 0
	// fmt.Println(slice)

	// 슬라이스 초기화
	// {}사용해서 초기화
	var slice1 = []int{1, 2, 3}
	var slice2 = []int{1, 5: 2, 10: 3}
	fmt.Println(slice1)
	// [1 2 3]
	fmt.Println(slice2)
	// [1 0 0 0 0 2 0 0 0 0 3]

	var arrayType = [...]int{1, 2, 3} // 배열: 고정길이
	var sliceType = []int{1, 2, 3}    // 슬라이스

	fmt.Println(arrayType)
	fmt.Printf("type of arrayType is %T\n", arrayType)
	// [1 2 3]
	// type of arrayType is [3]int
	fmt.Println(sliceType)
	fmt.Printf("type of sliceType is %T\n", sliceType)
	// [1 2 3]
	// type of sliceType is []int

	// make() 사용해서 초기화
	var sliceMake = make([]int, 3)
	fmt.Println(sliceMake)
	fmt.Printf("type of sliceMake is %T\n", sliceMake)
	// [0 0 0]
	// type of sliceMake is []int
}
