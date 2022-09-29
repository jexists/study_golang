package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var a, b int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for true {
		_, err := fmt.Fscanln(reader, &a, &b)
		if err != nil {
			break
		}
		fmt.Println(a + b)
	}

	// for true {
	// 	itemNumber, err := fmt.Fscanln(reader, &a, &b)
	// 	if itemNumber != 2 {
	// 		break
	// 	}
	// 	fmt.Println(a + b)
	// }
}

// 참고사이트
// https://jeonghwan-kim.github.io/dev/2019/01/08/go-fmt.html
