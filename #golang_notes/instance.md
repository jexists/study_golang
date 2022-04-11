## instance

→ Pointer타입의 struct다. (*Struct)

→ 생명 (객체의 생명체)

→ 주어

```go
package main

import "fmt"

type Student struct {
	name  string
	age   int
	grade int
}

func SetName(t Student, newName string) {
	t.name = newName
}

func main() {
	a := Student{"aaa", 20, 10}

	b := a
	b.age = 30

	// Value type
	fmt.Println(a) // {aaa 20 10}
	fmt.Println(b) // {aaa 30 10}

	c := Student{"aaa", 20, 10}

	var d *Student
	d = &c
	d.age = 30

	// Pointer (reference 타입형태 - 참조)
	fmt.Println(c) // {aaa 30 10}
	fmt.Println(d) // &{aaa 30 10}

	SetName(a, "bbb") //Value Type
	fmt.Println(a)    //{aaa 20 10}
}
```

## Function 형태

→ 기능 단위

→ procedure절차 (argument)

→ Input ⇒ 기능 순서 중요

```go
package main

import "fmt"

type Student struct {
	name  string
	age   int
	grade int
}

func SetName(t Student, newName string) {
	t.name = newName
}

//Function 형태 - 기능 단위, procedure절차 (Input), -> 기능 순서 중요
func SetNameFix(t *Student, newName string) {
	t.name = newName
}

func main() {
	a := Student{"aaa", 20, 10}

	SetNameFix(&a, "bbb") //Reference Type
	fmt.Println(a)        //{bbb 20 10}

}
```

## Method 형태

→ 매소드 형태의 표현 - object단위

→ OOP

→ 주체(object)/관계(relationship/주체(object))

→ object와 object의 관계 relationship 중요 (Entity-Relatioship)

→ 주어.관계(목적어)

→ 주어: Instance (생명주기)

```go
package main

import "fmt"

type Student struct {
	name  string
	age   int
	grade int
}

func SetName(t Student, newName string) {
	t.name = newName
}

//Function 형태 - 기능 단위, procedure절차 (Input), -> 기능 순서 중요
func SetNameFix(t *Student, newName string) {
	t.name = newName
}
//매소드 형태의 표현 - object단위
//(OOP, 주체(object)/관계(relationship/주체(object)) -> object와 object의 관계 relationship 중요 (Entity-Relatioship)
func (t *Student) SetNameFixM(newName string) {
	t.name = newName
}

func main() {
	a := Student{"aaa", 20, 10}

	a.SetNameFixM("bbb") //a=> instance
	fmt.Println(a) //{bbb 20 10}

}
```

### function vs method 결과는 똑같다.