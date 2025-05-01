package main

import "fmt"

// 알파벳 소문자로 이루어진 문자열 중 가장 많이 나오는 알파벳 출력
func main() {
	str := "sdlkjfadskafskdjfhweoijkvngdmfnlkfjwpefoj"
	var count [26]int
	for i := 0; i < len(str); i++ {
		count[str[i]-'a']++
	}

	maxCount := 0
	var maxCh byte
	for i := 0; i < 26; i++ {
		if count[i] > maxCount {
			maxCount = count[i]
			maxCh = byte('a' + i)
		}
	}

	fmt.Printf("max: %c count:%d", maxCh, maxCount)
}
