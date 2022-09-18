package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var num int
	fmt.Fscan(reader, &num)

	for i := 0; i < num; i++ {
		var a, b int
		fmt.Fscan(reader, &a, &b)
		fmt.Fprintln(writer, a+b)
	}
}
