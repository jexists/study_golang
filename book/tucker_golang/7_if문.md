
## if 문
- 만족하는 조건문의 {} 안에 있는 문장 실행
- 만족하는 조건문이 없으면 else 구문 {} 안에 있는 문장 실행
- else if & else 구문 생략가능

## 쇼트서킷 (short-circuit)
- &&의 좌변이 false이면 우변 무시 결과 false 처리
- ||의 좌변이 true이면 우변 무시 결과 true 처리

## if 초기문; 조건문
- if문 조건을 검사하기 전에 초기문에 입력 가능
- 검사에 사용할 변수를 초기화할때 사용
- 초기문 자리에 하나의 구문 가능, 끝에 ; 입력해서 구문이 끝남 표시 이후 조건문 입력

```go
package main

func main() {
	var 조건문 bool
	var 조건문1 bool
	var 조건문2 bool
	var 조건문3 bool

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
```