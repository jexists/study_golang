package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// 숫자 맞추기 게임
// - rand() 랜덤값 생성 패키지
// - rand.Intn() 0~N-1 사이의 값 생성

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}

func main() {

	// 랜덤한 숫자 생성
	randTest := rand.Intn(100)
	fmt.Println(randTest)
	// rand.Seed()가 0이라서 일정한 값 호출

	rand.Seed(time.Now().UnixNano())

	r := rand.Intn(100)
	cnt := 1

	// 숫자 맞추기 게임
	for {
		fmt.Printf("숫자값을 입력하세요: ")
		n, err := InputIntValue()
		if err != nil {
			fmt.Println("숫자만 입력하세요")
		} else {
			// fmt.Println("입력하신 숫자는 ", n, " 입니다.")
			// fmt.Println(r)
			if n > r {
				fmt.Println("입력하신 숫자가 더 큽니다.")
			} else if n < r {
				fmt.Println("입력하신 숫자가 더 작습니다.")
			} else {
				fmt.Println("숫자를 맞췄습니다. 축하합니다. 시도한 횟수: ", cnt)
				break
			}
			cnt++
		}
	}
}
