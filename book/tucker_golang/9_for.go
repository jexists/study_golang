package main

import "fmt"

func main() {

}

// 플래그 변수 사용
func flagFor() {
	a := 1
	b := 1
	found := false
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				found = true
				break
			}
		}
		if found {
			break
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}

// 레이블 사용
func lableFor() {
	a := 1
	b := 1
OutterFor:
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				break OutterFor
			}
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}

// 함수 사용
func find45(a int) (int, bool) {
	for b := 1; b <= 9; b++ {
		if a*b == 45 {
			return b, true
		}
	}
	return 0, false
}
func funcFor() {
	a := 1
	b := 0
	for ; a <= 9; a++ {
		var found bool
		if b, found = find45(a); found {
			break
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)

}
