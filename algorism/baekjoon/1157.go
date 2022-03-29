package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 알파벳 대소문자로 된 단어가 주어지면, 이 단어에서 가장 많이 사용된 알파벳이 무엇인지 알아내는 프로그램을 작성하시오. 단, 대문자와 소문자를 구분하지 않는다.

	var word string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &word)
	// fmt.Scanf("%s", &word)

	// 대문자로 변환
	word = strings.ToUpper(word)
	// fmt.Println(word)

	// array 생성 {아스키코드: 0}
	var letters = make(map[uint8]int)
	for i := 0; i < 26; i++ {
		letters[uint8(i)+65] = 0
	}
	// fmt.Println(letters)

	// 아스키코드++
	for i := 0; i < len(word); i++ {
		letters[word[i]]++
	}
	// fmt.Println(letters)

	var maxVal = -1
	var maxKey string
	for key, val := range letters {
		if val > maxVal {
			// 숫자 비교해서 maxVal maxKey 변경
			maxVal = val
			maxKey = string(key)
		} else if val == maxVal {
			// 같을경우 maxKey ? 변경
			maxKey = "?"
		}
	}

	fmt.Println(maxKey)
}
