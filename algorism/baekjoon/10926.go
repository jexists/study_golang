package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var a string
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &a)
	fmt.Println(a + "??!")
}
