package main

import "fmt"

type account struct {
	balance   int
	firstName string
	lastName  string
}

// 포인터 메서드
func (a1 *account) withdrawPointer(amount int) {
	a1.balance -= amount
}

// 값 메서드
func (a2 account) withdrawValue(amount int) {
	a2.balance -= amount
}

func (a3 account) withdrawReturnValue(amount int) account {
	a3.balance -= amount
	return a3
}

func main() {
	var mainA *account = &account{100, "joy", "jeong"}
	mainA.withdrawPointer(30)  // 같은 메모리
	fmt.Println(mainA.balance) // 70

	mainA.withdrawValue(20)    // 복사 (다른 메모리)
	fmt.Println(mainA.balance) // 70

	var mainB account = mainA.withdrawReturnValue(20) // 변경된 값 다시 반환
	fmt.Println(mainB.balance)                        // 50

	mainB.withdrawPointer(30)  // 같은 메모리
	fmt.Println(mainB.balance) // 20

}
