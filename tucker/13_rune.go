package main

import "fmt"

func main() {
	str := "Hello 월드"
	runes := []rune(str)
	// 총 바이트 길이
	fmt.Printf("len(str) = %d\n", len(str)) // len(str) = 12

	// 배열 요소 개수
	fmt.Printf("len(runes) = %d\n", len(runes)) // len(runes) = 8
}
