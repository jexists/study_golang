package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	stdin := bufio.NewReader(os.Stdin)
	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)

	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n') // 표준 입력 스트림 지우기
		// 입력받을때 에러 발생하면 표준 입력 스트림 삭제 필요
	} else {
		fmt.Println(n, a, b)
	}

	n, err = fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n, a, b)
	}

	// var c, d int
	// fmt.Scanln(&c, &d)
}
