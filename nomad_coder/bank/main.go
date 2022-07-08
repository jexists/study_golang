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
	fmt.Println(account)
}
