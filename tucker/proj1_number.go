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

// 전역변수 선언 (표준입력 스트림으로부터 값을 읽기)
var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	// 포인터로 넘겨줘야 값을 채워짐
	if err != nil {
		stdin.ReadString('\n')
		// fmt.Println(err)
		// 개행문자가 나올때까지 읽어서 비워야함 (입력오퍼에 남아있어서)
	}
	return n, err
}

func main() {

	// 랜덤한 숫자 생성 (0-99사이의 숫자)
	// randTest := rand.Intn(100)
	// fmt.Println("seedX :", randTest)

	// rand.Seed()가 0이라서 일정한 값 호출
	// 내부에 있는 시퀀스가 같아서 매번 호출할때마다 똑같음

	// 시드값(기반)을 주면 랜덤숫자 변경
	// 시드값도 똑같으면 값이 똑같음
	// 시드값이 바꿔져야 결과값이 바뀜
	// rand.Seed(time.Now().UnixNano())
	// rand.New()
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// 랜덤시드로 하면 좋은것? -> 매번 변화하는 값: Time
	// time.Now().UnixNano() : t의 1970년 1월 1일 부터 현재까지 경과된 시간을 nano sec단위로 반환

	r := rand.Intn(100)
	cnt := 1

	// fmt.Println("답 :", r)

	// 숫자 맞추기 게임
	for {
		fmt.Printf("숫자값을 입력하세요: ")
		n, err := InputIntValue()
		if err != nil {
			fmt.Println("숫자만 입력하세요.")
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
