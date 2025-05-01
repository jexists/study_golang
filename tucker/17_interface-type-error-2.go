// COMPILE ERROR: 문법에서부터 안됨
package main

type Stringer interface {
	String() string
}

type Student struct {
}

func main() {
	var stringer Stringer
	stringer.(*Student) // ERROR
	// impossible type assertion: stringer.(*Student)
	// Student 타입은 string() 메소드를 가지고 있지 않아서 생긴 에러 (불가능한 타입 변환: compile Error 빌드당시에 에러)
}

// ====================
// RUNTIME ERROR: 문법적으로는 가능
package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
}

func (s *Student) String() string {
	return "student"
}

func main() {
	var stringer Stringer
	s := stringer.(*Student) // ERROR
	fmt.Println(s)
	// panic: interface conversion: main.Stringer is nil, not *main.Student
}
