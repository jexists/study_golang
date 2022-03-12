# 변수 variable

## 변수

→ 다양하게 변할수 있는 것

→ 이름 , 값, type

→ 변수명 Type

→ 전역변수, 함수변수

→ 변수 = 메모리 번지(시작점), 사이즈(type)

→ 끝 = 주소 + 사이즈 

→ 변수: 이름, 값, 메모리주소, 사이즈

### 타입 종류

**정수**

→ int: 4/8byte 

→ int32: 4byte

→ int64: 8byte

→ int8: 1byte

→ int16: 2byte

**실수**

→ float32: 4byte

→ float64: 8byte

**bool: true/false**

**string: 문자열**

```python
package main

func main() {
	// 1. 선언하고 대입하는 방법
	// var a = int
	// a = 4
	// 2. 타입적지않고 쓰는방법
	// var a = 4
	// 3. 타입적고 대입
	// var a int = 4
	// 4. 한줄로 적는 방법
	a := 4
	b := 2

	fmt.printf("%v\n", a&b)
	fmt.printf("%v\n", a|b)
}
```

```go
package main

import "fmt"

// 전역변수
var c, python, java bool
var a int = 10

func main() {
	//함수변수
	var i int
	fmt.Println(i, c, python, java)
}
```

```python
package main

import "fmt"

func main() {
	//선언
	var a int //정수
	var b int

	//대입
	a = 3
	// a = 3.14 //float ERROR
	// a = "abc" //string ERROR
	b = 4

	fmt.Println(a + b) //7
}
```

### 변수 타입추론 / 초기화

→ 초깃값 존재하면 type 생략 가능

```go
package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
```

### 짧은 변수 선언

→ `변수명 := 변수값`

→ var 키워드 제외

→ 함수 밖에서는 모든 선언이 키워드(var, func 등) 시작해서 := 사용 불가 (전역변수는 무조건 var 키워드 사용!)

```go
package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3

	c := int32(1234)

	python, java := false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
```

## 상수

→ const

→ character, string, bollean, 숫자

→ := 사용불가

---

int8 → -128~127 (256개)

unit8 → 0~255

int16 → -32768~32767 (65535개)

uint16 → 0~65535

int32 → -21억~21억

unit32 → 42억