package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var a, b, c int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &a, &b, &c)
	fmt.Println((a + b) % c)
	fmt.Println(((a % c) + (b % c)) % c)
	fmt.Println((a * b) % c)
	fmt.Println(((a % c) * (b % c)) % c)

}
