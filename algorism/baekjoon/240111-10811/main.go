package main

import (
	"fmt"
	"slices"
)

// 12345
// 21345
// 21435
// 34125
// 34125

func main() {
	var n, m int // 바구니 수, 순서
	fmt.Scan(&n, &m)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}

	for i := 0; i < m; i++ {
		var first, last int
		fmt.Scan(&first, &last)
		var changeArr = arr[first-1 : last]
		slices.Reverse(changeArr)
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
}
