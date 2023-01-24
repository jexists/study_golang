package main

import "fmt"

func main() {
	// 0-10사이의 값을 갖는 배열을 정렬
	arr := []int{5, 1, 3, 2, 5, 2, 6, 8, 2, 0, 4, 5, 1, 6, 8, 2, 7, 9, 2, 1, 5, 6, 10}
	var count [11]int
	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}
	fmt.Println("count: ", count)
	// count:  [1 3 5 1 1 4 3 1 2 1 1]

	sorted := make([]int, 0, len(arr))
	for i := 0; i < 11; i++ {
		for j := 0; j < count[i]; j++ {
			sorted = append(sorted, i)
		}
	}
	fmt.Println("sorted: ", sorted)
	// sorted:  [0 1 1 1 2 2 2 2 2 3 4 5 5 5 5 6 6 6 7 8 8 9 10]
}
