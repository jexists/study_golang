package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	idx := 2
	// 삭제하는 다음 인덱스 이후부터 끝까지 복사
	for i := idx + 1; i < len(slice); i++ {
		slice[i-1] = slice[i]
	}
	// 복사 이후 마지막 삭제
	fmt.Println(slice) // [1 2 4 5 6 6]
	slice = slice[:len(slice)-1]
	fmt.Println(slice) // [1 2 4 5 6]

	slice1 := []int{1, 2, 3, 4, 5, 6}
	idx1 := 2
	// 위에 코드 한줄로
	slice1 = append(slice1[:idx1], slice1[idx1+1:]...)
	// 삭제할 이전 슬라이스 + 삭제할 이후 슬라이스
	fmt.Println(slice) // [1 2 4 5 6]

	// copy() 사용해서 삭제
	slice2 := []int{1, 2, 3, 4, 5, 6}
	idx2 := 2
	copy(slice2[idx2:], slice2[idx2+1:])
	fmt.Println(slice2) // [1 2 4 5 6 6]
	slice2 = slice2[:len(slice2)-1]
	fmt.Println(slice2) // [1 2 4 5 6]

}
