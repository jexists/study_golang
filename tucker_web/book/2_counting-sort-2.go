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

	for i := 1; i < 11; i++ {
		count[i] += count[i-1]
	}

	fmt.Println("count2: ", count)
	// count2:  [1 4 9 10 11 15 18 19 21 22 23]

	sorted := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		sorted[count[arr[i]]-1] = arr[i]
		count[arr[i]]--
	}

	fmt.Println("sorted: ", sorted)
	// sorted:  [0 1 1 1 2 2 2 2 2 3 4 5 5 5 5 6 6 6 7 8 8 9 10]

	// => 2N+K 알고리즘 (N+N+K)
	// 중복 for문 사용안해서 조금 더 효율적임
	// 중복숫자가 많은 경우 좋은 방식
}
