package main

import (
	"fmt"
)

func main() {
	// 불변(immutable): 일부만 변경 불가

	var str string = "Hello World"
	str = "핼로 월드"
	// str[2] = "N" // 일부 변경 불가 에러!

	str = "Hello World"
	fmt.Println(str) // Hello World
	var slice []byte = []byte(str)
	slice[2] = 'N'
	fmt.Println(str) // Hello World (변경안됨)
	// 슬라이스로 타입 변환시 문자열을 복사해서 새로운 메모리 공간을 만듬
	fmt.Printf("%s\n", slice) // HeNlo World

}
