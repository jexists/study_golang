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

	//
	fmt.Println("같은 배열을 가리키기 때문에 slice5를 변경할경우 6도 같이 변경")
	slice5 := make([]int, 3, 6)
	slice6 := append(slice5, 4, 5)
	fmt.Println("slice5: ", slice5, len(slice5), cap(slice5))
	fmt.Println("slice6: ", slice6, len(slice6), cap(slice6))
	// slice5:  [0 0 0] 3 6
	// slice6:  [0 0 0 4 5] 5 6

	slice5[1] = 100
	fmt.Println("slice5: ", slice5, len(slice5), cap(slice5))
	fmt.Println("slice6: ", slice6, len(slice6), cap(slice6))
	// slice5:  [0 100 0] 3 6
	// slice6:  [0 100 0 4 5] 5 6

	slice5 = append(slice5, 500)
	fmt.Println("slice5: ", slice5, len(slice5), cap(slice5))
	fmt.Println("slice6: ", slice6, len(slice6), cap(slice6))
	// slice5:  [0 100 0 500] 4 6
	// slice6:  [0 100 0 500 5] 5 6

	slice6 = append(slice6, 500)
	fmt.Println("slice5: ", slice5, len(slice5), cap(slice5))
	fmt.Println("slice6: ", slice6, len(slice6), cap(slice6))
	// slice5:  [0 100 0 500] 4 6
	// slice6:  [0 100 0 500 5 500] 6 6

	slice6[1] = 300
	fmt.Println("slice5: ", slice5, len(slice5), cap(slice5))
	fmt.Println("slice6: ", slice6, len(slice6), cap(slice6))
	// slice5:  [0 300 0 500] 4 5
	// slice6:  [0 300 0 500 5 500] 6 6

	// 빈공간이 없을 때 (cap) 새로운 기본배열의 2배 크기로 기본 배열의 요소와 새로운 배열 복사
	// cap: 새로운 배열의 길이 값(기존cap*2), len 기존 길이에 추가한 개수, 포인터 새로운 배열
	slice6 = append(slice6, 700)
	fmt.Println("slice5: ", slice5, len(slice5), cap(slice5))
	fmt.Println("slice6: ", slice6, len(slice6), cap(slice6))
	// slice5:  [0 300 0 500] 4 6
	// slice6:  [0 300 0 500 5 500 700] 7 12

	slice5[3] = 77
	fmt.Println("slice5: ", slice5, len(slice5), cap(slice5))
	fmt.Println("slice6: ", slice6, len(slice6), cap(slice6))
	// slice5:  [0 300 0 77] 4 6
	// slice6:  [0 300 0 500 5 500 700] 7 12
}
