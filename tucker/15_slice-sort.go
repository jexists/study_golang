package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}
func (s Students) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}
func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	slice := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(slice)   // 숫자 정렬시 사용 integer sort
	fmt.Println(slice) // [1 2 3 4 5 6]

	// # slice 정렬
	// int -> Ints()
	// float64 -> Float64s()
	// string -> strings()

	// 구조체 정렬
	s := []Student{
		{"A", 31},
		{"B", 52},
		{"C", 42},
		{"D", 38},
		{"E", 18},
	}

	// sort.Sort(s) // ERROR (interface 인자)
	// Len, Less, Swap 가지고 있는 타입만 Sort의 인자로 사용가능
	fmt.Println(s)
	// [{A 31} {B 52} {C 42} {D 38} {E 18}]
	sort.Sort(Students(s))
	fmt.Println(s)
	// [{E 18} {A 31} {D 38} {C 42} {B 52}]
}
