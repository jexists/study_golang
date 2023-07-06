package main

import (
	"fmt"
	"os"
)

func main() {
	// defer 명령문

	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("failed to create a file", err)
		return
	}

	defer fmt.Println("반드시 호출")
	defer f.Close()
	defer fmt.Println("파일 닫기")

	fmt.Println("파일에 Hello 적기")
	fmt.Fprintln(f, "Hello")

	// 파일에 Hello 적기
	// 파일 닫기
	// 반드시 호출
}
