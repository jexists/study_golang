# Defer

→ 둘러싼 함수가 종료할 때까지 어떠한 함수의 실행 연기

→ 연기된 호출의 인자는 즉시 평가되지만 그 함수 호출은 둘러싼 함수 졸요할때까지 수행X

→ 끝나는 시점에서 뒤에서부터 시작(역순 실행)

→ 함수호출을 지연시켜서 호출하기위한 키워드

```go
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
// hello
// world
```

```go
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
// counting
// done
// 9
// 8
// 7
// 6
// 5
// 4
// 3
// 2
// 1
// 0
```