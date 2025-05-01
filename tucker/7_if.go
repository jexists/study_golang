package main

import "fmt"

func main() {
	var 조건문 bool
	var 조건문1 bool
	var 조건문2 bool
	var 조건문3 bool

	if 조건문 {
		fmt.Println("문장") // 조건문이 true일때 실행
	}

	if 조건문 {
		fmt.Println("문장") // 조건문이 true일때 실행
	} else {
		fmt.Println("문장") // 조건문이 false일때 실행
	}

	if 조건문1 {
		fmt.Println("문장") // 조건문1이 true일때 실행
	} else if 조건문2 {
		fmt.Println("문장") // 조건문2이 true일때 실행
	} else if 조건문3 {
		fmt.Println("문장") // 조건문3이 true일때 실행
	} else {
		// else는 항상 맨 마지막 & 생략 가능
		fmt.Println("문장") // 모든 조건 false일때 실행
	}

	// 초기문 먼저 실행 후 조건 검사
	if 초기문 := function(); 조건문 {
		fmt.Println(초기문)
		fmt.Println("문장") // 조건문이 true일때 실행
	}
}

func function() bool {
	return true
}
