package main

import (
	"fmt"
)

func main() {

	var number int
	fmt.Scanf("%d", &number)

	count := 0
	init := number
	calc := 0
	l := true
	for l {
		count += 1
		calc = int(number/10) + (number % 10)
		number = (number % 10 * 10) + (calc % 10)
		if number == init {
			l = false
		}
	}
	fmt.Println(count)
}
