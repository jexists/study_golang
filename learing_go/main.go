package main

import "fmt"

func main() {

	var noSuggestNum = 1_2_3_4
	var suggestNum = 1_234

	fmt.Println(noSuggestNum) // 1234
	fmt.Println(suggestNum)   // 1234

	var interpretedString = "test \n \"hello\""
	var rawString = `test 
	"hello"`

	fmt.Println(interpretedString)
	fmt.Println(rawString)
}

// 해당 밑줄은 숫자 값 자체에 영향X
// → 숫자 맨앞이나 맨뒤에 넣을 수 없고 다른 숫자와 연결되게 사용X
// → 10진수 숫자를 천단위 나타낼때 추천 (예: 1_234)
// → 2진수,8진수,16진수를 1,2,4바이트 경계로 나누어 표시할때 추천
// 비추천) var num = 1_2_3_4
