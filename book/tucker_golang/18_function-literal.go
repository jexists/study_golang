package main

import "fmt"

type OpFn func(int, int) int

// 함수타입이 너무 길어져서 별칭타입을 만든다
func getOperatorType(op string) OpFn {
	if op == "+" {
		return func(a, b int) int {
			return a + b
		}
	} else if op == "*" {
		return func(a, b int) int {
			return a * b
		}
	} else {
		return nil
	}
}

func main() {
	var operator func(int, int) int
	operator = getOperator("+")
	var result = operator(3, 4)
	var resultSame = add(3, 4)
	fmt.Println(result)     // 7
	fmt.Println(resultSame) // 7
}
