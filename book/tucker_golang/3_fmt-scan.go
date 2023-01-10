package main

import (
	"fmt"
)

func main() {
	var a int
	var b int

	scan, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println(scan, err)
	} else {
		fmt.Println(scan, a, b)
	}

	var c int
	var d int
	scanf, err := fmt.Scanf("%d %d\n", &c, &d)
	if err != nil {
		fmt.Println(scanf, err)
	} else {
		fmt.Println(scanf, c, d)
	}

	var e int
	var f int
	scanln, err := fmt.Scanln(&e, &f)
	if err != nil {
		fmt.Println(scanln, err)
	} else {
		fmt.Println(scanln, e, f)
	}
	// 마지막 입력값 이후 엔터 키로 종료
}
