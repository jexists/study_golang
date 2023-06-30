package main

import "fmt"

// 빈 인터페이스 interface{}
// - 모든 타입이 가능

func PrintVal(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Printf("v is int %d\n", int(t))
	case float64:
		fmt.Printf("v is float64 %f\n", float64(t))
	case string:
		fmt.Printf("v is string %s\n", string(t))
	default:
		fmt.Printf("not supported type %T:%v\n", t, t)

	}
}

type Student struct {
	Age int
}

func main() {
	PrintVal(10)
	PrintVal(3.14)
	PrintVal("Hi")
	PrintVal(Student{15})
	// v is int 10
	// v is float64 3.140000
	// v is string Hi
	// not supported type main.Student:{15}
}
