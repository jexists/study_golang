# 함수

## 함수

→ 0개 또는 많은 인자 가능

→ 변수 이름 뒤 Type

→ 값: 함수의 인수나 반환 값 사용

→ 기능/모듈 (반복기능)

→ 모듈화/분리

→ func 함수명(인자 인자타입) 리턴타입 { return }

→ 함수 호출할때 값이 복사

```go
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```

→ 두개 이상 연속된 변수가 같은 type일 경우 마지막 빼고 생략 가능

예) `x int, y int` → `x, y int`

```go
package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```

→ 여러개 결과 반환 가능

```go
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
```

```go
package main

import "fmt"

func main() {
	a, b := fun1(2, 3)
	fmt.Println(a, b) //3 2
}

func fun1(x, y int) (int, int) {
	return y, x
}
```

### naked return

→ 인자가 없는 return 문: 주어진 반환 값 반환

→ 짧은 함수에서만 사용 (권장X)

```go
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return 
}

func main() {
	fmt.Println(split(17))
}
```

## 함수 클로저

→ Go 함수들은 클로저일 수도 있습니다. 

→ 클로저는 함수의 외부로부터 오는 변수를 참조하는 함수 값입니다. 

→ 함수는 참조된 변수에 접근하여 할당될 수 있습니다. 

→ 이러한 의미에서 함수는 변수에 "bound(바운드)" 됩니다. 그 예로, adder 함수는 클로저를 반환합니다. 

→ 각 클로저는 그 자체의 sum 변수에 bound(바운드) 되어 있습니다.

```go
package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

//0 0
//1 -2
//3 -6
//6 -12
//10 -20
//15 -30
//21 -42
//28 -56
//36 -72
//45 -90
```

### 재귀호출

→ 반복문 비슷

→ 함수호출 과정: 인자 기록  - 시작 포인트

→ 단계 느리다

→ 수학적 정의일 경우 사용

→ 함수형 언어

### 함수

→ 모듈화, 분리

→ 이름, Input(여러개, Type), Output(여러개, Type)