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

	fmt.Fscanln(reader, &a, &b)
	fmt.Println(float64(a) / float64(b))

}
