package main

import "fmt"

func main() {
	var year int
	fmt.Scanf("%d", &year)

	fmt.Println(year)
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}