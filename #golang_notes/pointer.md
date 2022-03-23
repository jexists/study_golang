# 포인터 (Pointers)

# 포인터

→ 값의 메모리 주소

→ 포인터 산술 지원X

→ 변수를 가르키는 것

변수 =  이름, 타입, 값, 메모리 주소

## * 연산자

→ `*Type`: type값을 가르키는 포인터 (zero value: nil)

→ 포인터가 가르키는 주소의 값

```go
var p *int

fmt.Println(*p) // 포인터 p를 통해 i 읽기
*p = 21         // 포인터 p를 통해 i 설정
```

## &연산자

→ 피연산자에 대한 포인터 생성

```go
i := 42
p = &i
```

```go
package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
// 42
// 21
// 73
```

### 포인터 호출

→ 메모리 주소만 복사

→ 포인터를 함수 인자로 받으면 메모리 주소만 복사

```go
package main

import "fmt"

type Student struct {
	name string
	age  int

	grade string
	class string
}

func (s Student) PrintGrade() {
	fmt.Println(s.class, s.grade)
}

func (s Student) InputGrade(class string, grade string) {
	s.class = class
	s.grade = grade
}

func main() {
	var s Student = Student{name: "Joy", age: 27, class: "math", grade: "A"}

 	s.InputGrade("English", "C")
 	s.PrintGrade() //math A
}
```

### 값 호출

→ 전체 값 복사

→ 값을 함수 인자로 받으면 전체 속성 복사