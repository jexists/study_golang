package main

import "fmt"

func main() {
	var length int
	fmt.Scan(&length)

	for i := 0; i < length; i++ {
		var s string
		fmt.Scan(&s)

		if len(s) == 1 {
			fmt.Println(s + s)
		} else {
			fmt.Println(string(s[0]) + string(s[len(s)-1]))
		}
	}
}
