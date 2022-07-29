package main

import (
	"bank/accounts"
	"fmt"
)

// func main() {
// 	account := banking.Account{Owner: "jexists", Balance: 1000}
// 	fmt.Println(account)
// }

// private로 변경후 struct making function 사용
func main() {
	account := accounts.NewAccount("jexists")
	account.Deposit(10)
	fmt.Println(account)
	fmt.Println(account.Balance())
	err := account.WithDraw(20)
	if err != nil {
		// Println() 호출 후 프로그램 종료
		// log.Fatalln(err)
		fmt.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account.String())
}
