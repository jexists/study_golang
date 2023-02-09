package main

import "fmt"

func changeArray(array2 [5]int) {
	array2[2] = 200
}

func changeSlice(slice2 []int) {
	slice2[2] = 200
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	// 배열의 모든값 복사
	changeArray(array)
	// 구조체 복사 (주소값 복사: 같은 주소값이여서 같은 데이터)
	changeSlice(slice)

	fmt.Println(array) // [1 2 3 4 5]
	fmt.Println(slice) // [1 2 200 4 5]

}
