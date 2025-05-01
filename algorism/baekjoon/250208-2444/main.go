package main

import "fmt"

// func main() {

// 	var n int
// 	fmt.Scan(&n)
// 	for i := 1; i <= n; i++ {
// 		for j := n - i; j > 0; j-- {
// 			fmt.Print(" ")
// 		}
// 		for k := 2*i - 1; k > 0; k-- {
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	}
// 	for i := n - 1; i > 0; i-- {
// 		for j := n - i; j > 0; j-- {
// 			fmt.Print(" ")
// 		}
// 		for k := 2*i - 1; k > 0; k-- {
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	}
// }

func main() {
	var n int
	fmt.Scan(&n)

	// 윗부분
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// 아랫부분
	for i := n - 2; i >= 0; i-- {
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
