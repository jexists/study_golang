package main

import "fmt"

func main() {
	type Person struct {
		ID     string
		Name   string
		Age    int
		Width  float64
		Height float64
	}

	// 초기값
	var person1 Person
	fmt.Println(person1)
	// {  0 0 0}

	var person2 Person = Person{"a", "joy", 28, 40.0, 160.0}
	fmt.Println(person2)
	// {a joy 28 40 160}

	// 여러줄로 초기화할 때 쉼표 필수
	var person3 Person = Person{
		"a",
		"joy",
		28,
		40.0,
		160.0,
	}
	fmt.Println(person3)
	// {a joy 28 40 160}

	var person4 Person = Person{
		Age:    28,
		Height: 160.0,
	}
	fmt.Println(person4)
	// {  28 0 160}

	// 구조체 + 구조체
	type student struct {
		Person Person
		No     int
		ID     string
	}

	var student1 student = student{
		Person{"a", "joy", 28, 40.0, 160.0},
		1,
		"b",
	}
	fmt.Println(student1)
	// {{a joy 28 40 160} 1 b}

	fmt.Println(student1.ID)          // a
	fmt.Println(student1.No)          // 1
	fmt.Println(student1.Person)      // {a joy 28 40 160}
	fmt.Println(student1.Person.ID)   // a
	fmt.Println(student1.Person.Name) // joy
	// fmt.Println(Person.Name)       // ERROR

	// 포함된 필드: 구조체 안에 포함된 다른 구조체의 필드명을 생략하는 경우
	type studentVIP struct {
		Person
		No int
		ID string
	}

	var student2 studentVIP = studentVIP{
		Person{"c", "sad", 29, 50.0, 170.0},
		2,
		"d",
	}
	fmt.Println(student2)
	// {{c sad 29 50 170} 2 d}

	// 포함된 필드: 구조체 안에 포함된 다른 구조체의 필드명을 생략하는 경우
	fmt.Println(student2.ID)          // d
	fmt.Println(student2.No)          // 2
	fmt.Println(student2.Person)      // {c sad 29 50 170}
	fmt.Println(student2.Person.ID)   // c
	fmt.Println(student2.Person.Name) // sad
	fmt.Println(student2.Name)        // sad

}
