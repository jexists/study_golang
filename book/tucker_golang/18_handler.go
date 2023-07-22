package main

import (
	"fmt"
	"os"
)

type Writer func(string)

func writeHello(writer Writer) {
	writer("hello world")
}

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("fail")
		return
	}
	defer f.Close()

	writeHello(func(msg string) {
		fmt.Fprintln(f, msg)
	})
}
