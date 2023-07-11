package main

import "fmt"

// func mainErr() {
// 	i += 10 // 불가능
// }

func main() {
	i := 0
	// 실행된 상태가 아님 (위치)
	f := func() {
		i += 10
		// 내부 상태를 가져올 수 있음
		// 값복사가 아닌 레퍼런스 복사 (포인터 복사)
		// i := &i (포인터 복사)
	}
	i++
	fmt.Println(i) // 1
	// 함수 실행
	f()
	fmt.Println(i) // 11
}
