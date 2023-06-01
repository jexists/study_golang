package main

import "fmt"

func main() {

	// 복제: 같은 길이 슬라이스 생성 후 순회해서 요솟값 복사
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))
	for i, v := range slice1 {
		slice2[i] = v
	}

	slice1[1] = 100

	fmt.Println(slice1)
	fmt.Println(slice2)

	// append() 함수 사용
	slice := []int{1, 2, 3, 4, 5}
	slice3 := append([]int{}, slice...)
	fmt.Println(slice3)
	// [1 2 3 4 5]

	slice4 := append([]int{}, slice[0], slice[1], slice[2], slice[3], slice[4])
	fmt.Println(slice4)
	// [1 2 3 4 5]

	// copy()
	// func copy(dst, src []Type) int
	// func copy(복사한 결과를 저장하는 슬라이스 변수, 복사대상이 되는 슬라이스) 복사된 요소 개수(길이가 작은 갯수만큼 복사)

	slice5 := []int{1, 2, 3, 4, 5}
	slice6 := make([]int, 3, 10)
	slice7 := make([]int, 10)
	cnt1 := copy(slice6, slice5)
	cnt2 := copy(slice7, slice5)

	fmt.Println(len(slice5), cap(slice5))
	fmt.Println(len(slice6), cap(slice6))
	fmt.Println(len(slice7), cap(slice7))
	// 5 5
	// 3 10
	// 10 10
	fmt.Println(cnt1, slice6)
	fmt.Println(cnt2, slice7)
	// 3 [1 2 3]
	// 5 [1 2 3 4 5 0 0 0 0 0]

	slice8 := make([]int, len(slice5))
	// 길이가 같은 슬라이스 만든 후 복사
	cnt3 := copy(slice8, slice5)
	fmt.Println(len(slice8), cap(slice8))
	fmt.Println(cnt3, slice8)
	// 5 5
	// 5 [1 2 3 4 5]
}
