# 반복문

# For

→ 하나의 반복 구조

→ `for 초기화구문; 조건표현; 사후 구문 {}`

- 초기화구문: 첫번째 iteration 전 수행 (초기화 구문에 선언된 변수는 for문에서만 사용 / 옵션)
- 조건 표현: 모든 iteration 이전 판별 (false일 경우 iterating 종료)
- 사후 구문: iteration 마지막 수행 (옵션)

→ {} 필수

→ 조건문이 true인 한 계속 수행

```go
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum) // 45
}
```

→ 초기화 구문, 사후 구문 필수 x

```go
package main

import "fmt"

func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum) //1024
}
```

**break: 종료**

```go
func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
}
//0
//1
//2
//3
//4
```

**continue: 그 다음걸로 이동**

```go
func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
//0
//1
//2
//3
//4
//6
//7
//8
//9
```

→ while (for문에 ; 생략하면 비슷)

```go
package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum) //1024
}
```

### 무한루프

→ 모두 생략할때 무한루프

→ 반복조건 생략 하면 무한 루프

```go
package main

func main() {
	for {
	}
}
```

```go
//BUG 계속 돌아간다
package main

func main() {
	var i int
	for {
		if i == 5 {
			continue
		}
		if i == 6 {
			break
		}
		fmt.Println(i)
	i++
	}
}
//0
//1
//2
//3
//4
//계속실행
```

```go
package main

func main() {
	var i int
	for {
		if i == 5 {
			i++
			continue
		}
		if i == 6 {
			break
		}
		fmt.Println(i)
	i++
	}
}
//0
//1
//2
//3
//4
```

## 변수의 scope

```go
func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("최종 i 값", i) //ERROR
}
```

```go
// 수정값
func main() {
	var i int
	for i = 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("최종 i 값", i) // 5
}
```

```go
// i가 다르다
func main() {
	var i int
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("최종 i 값", i) // 0
}
```


```go
package main

import "fmt"

func superAdd(numbers ...int) int {
	fmt.Println(numbers)
	// for each, for in == for range
	for index, number := range numbers {
		fmt.Println(index, number)
		total += number
	}
	return total
}

func main() {
	total := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(total)
}
```
