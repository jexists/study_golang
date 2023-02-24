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
		fmt.Println(a - b)
	}

}
