package main

import "fmt"

func main() {
	// 타입 변환 ========================
	a := 3
	var b float64 = 3.5

	var c int = int(b)  // 3
	d := float64(a * c) // 9 := 3 * 3

	var e int64 = 7
	f := int64(d) * e // 10 := 9 * 7

	var g int = int(b * 3) // 10 := (3.5 * 3) -> (10.5) -> 10
	// 실수 -> 정수 변환시 내림 (소수점 이하 숫자 반올림X, 내림)
	var h int = int(b) * 3

	fmt.Println(g, h, f) // 10 9 63

}
