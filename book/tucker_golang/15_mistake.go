package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3}
	addNum(slice)
	fmt.Println(slice) // [1, 2, 3]
	// 포인터를 넘기니깐 바뀔거라고 생각할수도 있지만..
	// 서로 다른 인스턴스라서 안바뀜

	// 기존 배열 변경방법 - 포인터 사용
	addNumPointer(&slice)
	fmt.Println(slice) // [1, 2, 3, 4]

	addNum(slice)
	fmt.Println(slice) // [1, 2, 3, 4]

	// 기존 배열 변경방법 - 슬라이스 리턴
	slice = addNumReturn(slice)
	fmt.Println(slice) // [1, 2, 3, 4, 4]

}

func addNum(slice []int) {
	slice = append(slice, 4)
}

func addNumPointer(slice *[]int) {
	// 포인터 (copy되는 데이터가 작긴한데 헷갈릴수있음)
	*slice = append(*slice, 4)
}

func addNumReturn(slice []int) []int {
	// 새로운 슬라이스를 값을 넘긴다.
	slice = append(slice, 4)
	return slice
}
