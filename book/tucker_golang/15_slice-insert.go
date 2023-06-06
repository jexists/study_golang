package main

import "fmt"

func main() {
	// 요소 추가
	slice := []int{1, 2, 3, 4, 5, 6}
	slice = append(slice, 0)
	idx := 2
	for i := len(slice) - 2; i >= idx; i-- {
		slice[i+1] = slice[i]
	}
	fmt.Println(slice) // [1 2 3 3 4 5 6]
	slice[idx] = 100
	fmt.Println(slice) // [1 2 100 3 4 5 6]

	// 중간에 메모리(벙커)를 더 사용 (효율성 떨어짐)
	slice2 := []int{1, 2, 3, 4, 5, 6}
	idx2 := 2
	slice2 = append(slice2[:idx2], append([]int{100}, slice2[idx2:]...)...)
	fmt.Println(slice2) // [1 2 100 3 4 5 6]

	// 불필요한 메모리 할당없이 (copy() 사용)
	slice3 := []int{1, 2, 3, 4, 5, 6}
	idx3 := 2
	slice3 = append(slice3, 0)
	copy(slice3[idx3+1:], slice3[idx3:])
	fmt.Println(slice3) // [1 2 3 3 4 5 6]
	slice3[idx3] = 100
	fmt.Println(slice3) // [1 2 100 3 4 5 6]
}
