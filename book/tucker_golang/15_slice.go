package main

import "fmt"

func main() {
	var array [10]int
	fmt.Println(array)
	// [0 0 0 0 0 0 0 0 0 0]

	var slice []int
	// 슬라이스일 경우 []대괄호 안에 길이 넣지 말아야함
	if len(slice) == 0 {
		fmt.Println("slice empty ", slice)
		// slice empty  []
	}
	// slice[1] = 10 // ERROR
	// panic: runtime error: index out of range [1] with length 0
	// fmt.Println(slice)
	// 초기화 할경우 길이가 0

	// 슬라이스 초기화
	// {}사용해서 초기화
	var slice1 = []int{1, 2, 3}
	var slice2 = []int{1, 5: 2, 10: 3}
	fmt.Println(slice1)
	// [1 2 3]
	fmt.Println(slice2)
	// [1 0 0 0 0 2 0 0 0 0 3]

	var arrayType = [...]int{1, 2, 3} // 배열: 고정길이
	// [...]int{1, 2, 3} == [3]int{1,2,3}
	var sliceType = []int{1, 2, 3} // 슬라이스

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
	fmt.Println(sliceMake) // [0 0 0]
	fmt.Printf("type of sliceMake is %T\n", sliceMake)
	// type of sliceMake is []int
	fmt.Println(len(sliceMake)) // 3
	fmt.Println(cap(sliceMake)) // 3

	var sliceMakeWithCap = make([]int, 3, 5)
	fmt.Println(sliceMakeWithCap)      // [0 0 0]
	fmt.Println(len(sliceMakeWithCap)) // 3
	fmt.Println(cap(sliceMakeWithCap)) // 5
}
