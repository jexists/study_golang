package main

import "fmt"

func changeArray(array2 [5]int) {
	// 새로운 값
	array2[2] = 200
}

func changeSlice(slice2 []int) {
	// 포인터
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
	// rvalue (값)
	// array != array2 (대입)
	// 대입: 공간 = 값
	// 대입연산자 일경우 좌변 우변 타입이 같아야한다.
	fmt.Println(slice) // [1 2 200 4 5]

	array2 := [100]int{1: 1, 2: 2, 99: 100}
	slice2 := array2[1:10]
	slice3 := slice2[2:99]
	fmt.Println(slice2)
	// [1 2 0 0 0 0 0 0 0]
	fmt.Println(slice3) // 가르키는 배열을 보고잇음
	// [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 100]
}
