package main

import "fmt"

func main() {
	// 배열 변수 선언 및 초기화

	// 초깃값 지정하지 않을 경우 기본값으로 초기화
	var nums [5]int
	fmt.Println(nums)
	// [0 0 0 0 0]

	// 초깃값 지정
	days := [3]string{"monday", "tuesday", "wednesday"}
	fmt.Println(days)
	// [monday tuesday wednesday]

	// 지정한 초깃값만 초기화 되고 나머지는 기본값으로 초기화
	var temps [5]float64 = [5]float64{24.3, 26.7}
	fmt.Println(temps)
	// [24.3 26.7 0 0 0]

	// 인덱스로 지청한 값만 초기화
	var s = [5]int{1: 10, 3: 30}
	fmt.Println(s)
	// [0 10 0 30 0]

	// 배열요소 갯수 생략 가능 (배열요소 개수는 초기화 되는 요소개수와 같음)
	x := [...]int{10, 20, 30}
	fmt.Println(x)
	// [10 20 30]

	y := [...]int{}
	fmt.Println(y)
	// []

}
