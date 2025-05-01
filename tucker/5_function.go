package main

import "fmt"

func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)
	// 3 true
}

// 멀티반환함수: 여러개 반환가능 / 반환 타입을 소괄호() 묶어서 표현
// 반환할 변수 변수명 지정해서 반환시 return에 해당 변수 안적어도됨
func Divide(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return
	}
	result = a / b
	success = true
	return
}
