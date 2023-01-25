package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 숫자 맞추기 게임

	// 랜덤한 숫자 생성
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(100)
	fmt.Println(n)
}
