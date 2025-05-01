package main

import "fmt"

const PI = 3.14
const FloatPI float64 = 3.14

// 타입없는 상수를 선언시 타입이 다른 여러 변수에서 사용 가능
func main() {
	// var a int = PI // ERROR
	var b int = PI * 100
	// var c int = FloatPI * 100 // Error

	// fmt.Println(a)
	fmt.Println(b)
	// fmt.Println(c)
}
