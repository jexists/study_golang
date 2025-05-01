package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	// 실수 비교시 에러
	fmt.Printf("%f+%f == %f : %v\n", a, b, c, a+b == c)
	// 0.100000+0.200000 == 0.300000 : false
	fmt.Println(a + b)
	// 0.30000000000000004
	fmt.Println(c)
	// 0.3

	// 1. 매우 작은 값을 선정해서 오차 무시하는 방법
	fmt.Printf("%v / %v\n", a+b == c, equal(a+b, c))
	// false / true

	// 2. math.Nextafter()함수 사용
	fmt.Printf("%v / %v\n", a+b == c, equalUsingMath(a+b, c))
	// false / true

	// useMathBig
	useMathBig()
}

// 1. 매우 작은 값을 선정해서 오차 무시하는 방법
const epsilon = 0.000001 // 작은 값
func equal(a, b float64) bool {
	if a > b {
		if a-b <= epsilon {
			return true
		} else {
			return false
		}
	} else {
		if b-a <= epsilon {
			return true
		} else {
			return false
		}
	}
}

// 2. math.Nextafter()함수 사용
func equalUsingMath(a, b float64) bool {
	return math.Nextafter(a, b) == b
}

// 3. math/Big float 사용하는 방법
func useMathBig() {
	d, _ := new(big.Float).SetString("0.1")
	e, _ := new(big.Float).SetString("0.2")
	f, _ := new(big.Float).SetString("0.3")

	g := new(big.Float).Add(d, e)

	fmt.Println(d, e, f, g)
	// 0.1 0.2 0.3 0.3

	fmt.Println(f.Cmp(g)) // 0
	// f.Cmp(g) 함수: f == g 값비교
	// 리턴값: -1 (f < g)
	// 리턴값: 1 (f > g)
	// 리턴값: 0 (f == g)

}
