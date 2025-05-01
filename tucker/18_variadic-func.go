package main

import "fmt"

func numPrint(nums ...int) int {
	sum := 0
	fmt.Printf("nums 타입 %T\n", nums)
	for _, v := range nums {
		sum += v
	}
	return sum
}

// 가변인자받는거랑 슬라이스 받는거랑 같음
func numPrintSlice(nums []int) int {
	sum := 0
	fmt.Printf("nums 타입 %T\n", nums)
	for _, v := range nums {
		sum += v
	}
	return sum
}

func stringPrint(words ...string) string {
	sentence := ""
	fmt.Printf("words 타입 %T\n", words)
	for _, v := range words {
		sentence += v
	}
	return sentence
}

func anyPrint(any ...interface{}) string {
	fmt.Printf("words 타입 %T\n", any)
	val := ""
	for _, arg := range any {
		switch arg.(type) {
		case bool:
			// val = arg.(bool)
			val += fmt.Sprintf("%b, ", arg)
		case float64:
			// val = arg.(float64)
			val += fmt.Sprintf("%f, ", arg)
		case int:
			val += fmt.Sprintf("%d, ", arg)
			// val = arg.(int)
		case string:
			val += fmt.Sprintf("%s, ", arg)
		}
	}
	return val
}

func main() {
	fmt.Println(numPrint(1, 2, 3))
	// nums 타입 []int
	// 6
	fmt.Println(numPrint(1))
	// nums 타입 []int
	// 1
	fmt.Println(numPrint())
	// nums 타입 []int
	// 0

	fmt.Println(numPrintSlice([]int{1, 2, 3}))
	// nums 타입 []int
	// 6
	fmt.Println(numPrintSlice([]int{1}))
	// nums 타입 []int
	// 1
	fmt.Println(numPrintSlice([]int{}))
	// nums 타입 []int
	// 0

	fmt.Println(stringPrint("a", "b", "c"))
	// words 타입 []string
	// abc

	fmt.Println(anyPrint(1, "h", 3.14))
	// words 타입 []interface {}
	// 1, h, 3.140000,
}
