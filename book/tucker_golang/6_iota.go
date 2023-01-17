package main

import "fmt"

// 1, 2, 3.. 처럼 1씩 증가하도록 정의할 때 사용
// 소괄호를 벗어나면 다시 초기화
func main() {
	const (
		Red   int = iota // 0
		Blue  int = iota // 1
		Green int = iota // 2
	)
	fmt.Println(Red)   // 0
	fmt.Println(Blue)  // 1
	fmt.Println(Green) // 2

	const (
		C1 uint = iota + 1 // 1 = 0 + 1
		C2                 // 2 = 1 + 1
		C3                 // 3 = 2 + 1
	)
	fmt.Println(C1) // 1
	fmt.Println(C2) // 2
	fmt.Println(C3) // 3
}
