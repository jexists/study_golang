package main

import "fmt"

func main() {
	// 배열 스라이싱
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:2]
	fmt.Println("array: ", array, len(array), cap(array))
	fmt.Println("slice: ", slice, len(slice), cap(slice))
	// array:  [1 2 3 4 5] 5 5
	// slice:  [2] 1 4

	// 같이 변경
	array[1] = 100
	fmt.Println("array: ", array)
	fmt.Println("slice: ", slice, len(slice), cap(slice))
	// array:  [1 100 3 4 5]
	// slice:  [100] 1 4

	slice = append(slice, 500)
	fmt.Println("array: ", array)
	fmt.Println("slice: ", slice, len(slice), cap(slice))
	// array:  [1 100 500 4 5]
	// slice:  [100 500] 2 4

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1[1:2]
	fmt.Println("slice1: ", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2: ", slice2, len(slice2), cap(slice2))
	// slice1:  [1 2 3 4 5] 5 5
	// slice2:  [2] 1 4

	// 처음부터 슬라이싱
	slice4 := []int{1, 2, 3, 4, 5}
	slice5 := slice4[0:3]
	slice6 := slice4[:3]
	fmt.Println("slice4: ", slice4, len(slice4), cap(slice4))
	fmt.Println("slice5: ", slice5, len(slice5), cap(slice5))
	fmt.Println("slice6: ", slice6, len(slice6), cap(slice6))
	// slice4:  [1 2 3 4 5] 5 5
	// slice5:  [1 2 3] 3 3
	// slice6:  [1 2 3] 3 3

	// 끝까지 슬라이싱
	slice7 := []int{1, 2, 3, 4, 5}
	slice8 := slice7[2:len(slice7)]
	slice9 := slice7[2:]
	fmt.Println("slice7: ", slice7, len(slice7), cap(slice7))
	fmt.Println("slice8: ", slice8, len(slice8), cap(slice8))
	fmt.Println("slice9: ", slice9, len(slice9), cap(slice9))
	// slice7:  [1 2 3 4 5] 5 5
	// slice8:  [3 4 5] 3 3
	// slice9:  [3 4 5] 3 3

	// 전체 슬라이싱
	slice10 := []int{1, 2, 3, 4, 5}
	slice11 := slice10[0:len(slice10)]
	slice12 := slice10[:]
	fmt.Println("slice10: ", slice10, len(slice10), cap(slice10))
	fmt.Println("slice11: ", slice11, len(slice11), cap(slice11))
	fmt.Println("slice12: ", slice12, len(slice12), cap(slice12))
	// slice10:  [1 2 3 4 5] 5 5
	// slice11:  [1 2 3 4 5] 5 5
	// slice12:  [1 2 3 4 5] 5 5

	// cap 크기 조절
	slice13 := []int{1, 2, 3, 4, 5}
	slice14 := slice13[1:3]
	slice15 := slice13[1:3:4]
	fmt.Println("slice13: ", slice13, len(slice13), cap(slice13))
	fmt.Println("slice14: ", slice14, len(slice14), cap(slice14))
	fmt.Println("slice15: ", slice15, len(slice15), cap(slice15))
	// slice13:  [1 2 3 4 5] 5 5
	// slice14:  [2 3] 2 4
	// slice15:  [2 3] 2 3
	slice16 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	slice17 := slice16[1:3]
	slice18 := slice16[1:3:4]
	fmt.Println("slice16: ", slice16, len(slice16), cap(slice16))
	fmt.Println("slice17: ", slice17, len(slice17), cap(slice17))
	fmt.Println("slice18: ", slice18, len(slice18), cap(slice18))
	// slice13:  [1 2 3 4 5 6 7 8] 8 8
	// slice14:  [2 3] 2 7
	// slice15:  [2 3] 2 3
}
