package main

import "fmt"

func main() {
	var 비교값 int
	var 값1 int
	var 값2 int

	switch 비교값 { // 검사하는 값
	case 값1: // 비교값이 값1이 같을 때 수행
		fmt.Println("문장")
	case 값2: // 비교값이 값2이 같을 때 수행
		fmt.Println("문장")
	default: // 만족하는 case가 없을 때 수행
		fmt.Println("문장")
	}

	var day string
	// 하나의 case 한번에 여러값 비교
	switch day {
	case "monday", "tuesday", "wednesday", "thursday", "friday":
		fmt.Println("평일")
	case "saturday", "sunday":
		fmt.Println("주말")
	}

	// 조건문 비교
	// - if문처럼 true가 되는 조건문 검사
	temp := 18
	switch true {
	case temp < 10, temp < 30:
		fmt.Println("추운날씨")
	case temp > 15, temp < 30:
		fmt.Println("약간 추위")
	case temp >= 15, temp < 25:
		// 실행안함 (두번째 조건에서 실행후 종료)
		fmt.Println("좋은 날씨")
	default:
		fmt.Println("따뜻")
	}
}
