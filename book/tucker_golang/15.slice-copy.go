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
}
