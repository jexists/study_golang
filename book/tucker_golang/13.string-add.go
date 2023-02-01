package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func main() {
	// string 합 연산시 새로운 메모리 공간 만들어서 두 문자열 합침 (주솟값 변경 -> 메모리 낭비)
	str1 := "Hello"
	str2 := "월드"
	str3 := str1 + " " + str2
	fmt.Println(str3) // Hello 월드

	str1 += " " + str2
	fmt.Println(str1) // Hello 월드

	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	fmt.Println(stringHeader1) // &{824634425376 12}

	// strings.Builder사용
	str := "Hello World"
	fmt.Println(ToUpper(str)) // HELLO WORLD
}

func ToUpper(str string) string {
	var builder strings.Builder
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			builder.WriteRune('A' + (c - 'a'))
		} else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}
