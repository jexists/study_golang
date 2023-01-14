package main

import "fmt"

func main() {
	printNo(3)
	// main() -> printNo(3) -> printNo(2) -> printNo(1) -> printNo(0) -> 탈출 -> 자신 호출한 위치로 차례대로 연속해서 다시 돌아감 -> 최초 호출 main()함수에 도착
	// 3
	// 2
	// 1
	// After 1
	// After 2
	// After 3
}

func printNo(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	printNo(n - 1)
	fmt.Println("After", n)
}
