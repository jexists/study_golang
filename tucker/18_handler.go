package main

import (
	"fmt"
	"os"
)

type Writer func(string)

// 의존성 주입
func writeHello(writer Writer) {
	// 화면에 쓰는지 터미널에 쓰는지 모른다
	writer("hello world")
}

// -----

// 인터페이스
type WriterInterface interface {
	Write(string)
}

func writeHello2(writer WriterInterface) {
	writer.Write("hello world Interface")
}

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("fail")
		return
	}
	defer f.Close()

	// 파일에 hello world 작성
	// writeHello(func(msg string) {
	// 	fmt.Fprintln(f, msg)
	// })

	// 화면에 hello world 출력
	writeHello(func(msg string) {
		fmt.Println(msg)
	})
}
