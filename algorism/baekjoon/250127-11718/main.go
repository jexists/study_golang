package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// bufio.NewScanner(os.Stdin)
// -> 표준 입력을 스캔 (여러줄)

// scanner.Scan()
// -> 한 줄씩 입력을 읽음
// -> 입력이 없을 때 false 반환

// scanner.Text()
// -> 읽은 한 줄의 텍스트를 반환
