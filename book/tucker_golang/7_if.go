package main

func main() {
	var 조건문 bool
	var 조건문1 bool
	var 조건문2 bool
	var 조건문3 bool
	var 문장 func() void

	if 조건문 {
		문장 // 조건문이 true일때 수행
	}

	if 조건문 {
		문장 // 조건문이 true일때 수행
	} else {
		문장 // 조건문이 false일때 수행
	}

	if 조건문1 {
		문장 // 조건문1이 true일때 수행
	} else if 조건문2 {
		문장 // 조건문2이 true일때 수행
	} else if 조건문3 {
		문장 // 조건문3이 true일때 수행
	} else {
		// else는 항상 맨 마지막 & 생략 가능
		문장 // 모든 조건 false일때 수행
	}

	// 초기문 먼저 실행 후 조건 검사
	if 초기문; 조건문 {
		문장 // 조건문이 true일때 수행
	}
}
