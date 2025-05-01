package main

import "fmt"

type myInt int

func (a myInt) add(b int) int {
	// return a + b // ERROR
	return int(a) + b
}

func main() {
	var a myInt = 10
	fmt.Println(a.add(30))
	var b int = 20
	// fmt.Println(b.add(50))  // ERROR
	// 다른 타입이라서 메서드 사용불가
	fmt.Println(myInt(b).add(50))
}
