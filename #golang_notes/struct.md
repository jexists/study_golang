# 구조체

→ structure

→ 응집성: cohesive⬆️ / 종속성: defendency ⬇️

→ 객체

→ 속성 + 기능

→ OOP 기능 (=class)

# Structs

→ 필드의 집합체

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X) //4
}
```

### 구조체 포인터 통해서 구조체 필드 접근

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9 //(*p).X 같은 표현
	fmt.Println(v) //{1000000000 2}
}
```

### Struct Literals

→ 필드 값을 나열하여 새로 할당된 구조체 값 나타냄

→ & 구조체 값 포인터를 반환

```go
package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, v2, v3, p) //{1 2} {1 0} {0 0} &{1 2}
}
```

struct 기능

`func (변수 스트럭트) 기능이름() {}`

```go
package main

import "fmt"

// struct 속성(statement)/타입
type Student struct {
	name  string
	class int

	grade Grade
}

type Grade struct {
	name  string
	grade string
}

//Student struct가 가지고 있는 기능 (객체에 속한 매소드)
func (s Student) ViewGrade() {
	fmt.Println(s.grade)
}

func (s Student) InputGrade(name string, grade string) {
	s.grade.name = name
	s.grade.grade = grade
}

//일반함수
func ViewGrade(s Student) {
	fmt.Println(s.grade)
}

func InputGrade(s Student, name string, grade string) {
	s.grade.name = name
	s.grade.grade = grade
}

func main() {

	var s Student
	s.name = "joy"
	s.class = 1

	s.grade.name = "Math"
	s.grade.grade = "C"

	s.ViewGrade() //{Math C}
	ViewGrade(s)  //{Math C}

	s.InputGrade("English", "A")
	s.ViewGrade() //{Math C} *고랭 함수호출은 복사 (입력값 복사)
}
```


```go
package main

import "fmt"

type person struct {
	name string
	age int
	favFood []string
}

func main() {
	favFood := []string{"apple"}
	// struct 순서대로 입력하기
	jexist := person{"jexist", 12, favFood}

	fmt.Println(jexist)
	// struct key 적기
	jexist := person{name:"jexist", age:12, favFood: favFood}
	fmt.Println(jexist)

	// jexist := person{name:"jexist", 12, favFood} //ERROR
}
```