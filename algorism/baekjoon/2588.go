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
	fmt.Println((b / 100))
	fmt.Println((b / 10))
	// fmt.Println(a * (b % 10))
	// fmt.Println(((a % c) + (b % c)) % c)
	// fmt.Println((a * b) % c)
	// fmt.Println(((a % c) * (b % c)) % c)

}
