package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("안녕! 나는 %d살 %s이야", s.Age, s.Name)
}

func (s Student) GetAge() int {
	return s.Age
}
func main() {
	student := Student{"Joy", 12}
	var stringer Stringer
	stringer = student
	fmt.Printf("%s\n", stringer.String())
	// fmt.Printf("%d\n", stringer.GetAge()) // ERROR

}
