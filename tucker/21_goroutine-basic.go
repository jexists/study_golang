package main

import (
	"fmt"
	"time"
)

func PrintHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", v)
	}
}

func PrintNumbers() {
	for i := 0; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func main() {
	go PrintHangul()
	go PrintNumbers()

	// time.Sleep(3 * time.Second)
	// 종료

	// time.Sleep(3 * time.Second)
	// 가 0 나 1 다 2 라 마 3 바 4 사 5 %

	time.Sleep(1 * time.Second)
	// 메인함수가 종료되면 프로그램이 끝남
	// 가 0 나 1 다 %
}
