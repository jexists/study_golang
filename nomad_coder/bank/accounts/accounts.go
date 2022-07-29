package accounts

import (
	"errors"
	"fmt"
)

// export해서 사용해야하기때문에 첫글자 대문자
// private owner / public Owner
// type Account struct {
// 	Owner   string
// 	Balance int
// }

// Account struct
type Account struct {
	owner   string
	balance int
}

// 에러 생성시 err+에러명
var errNoMoney = errors.New("돈부족!")

// struct making function
// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
// func (a Account) Deposit(amount int) {
//  복사본을 만들어서 준 값을 그대로 안줌
//  복사본 만들지 말고 호출한 값을 사용해아함
// 	fmt.Println(amount) // 10
// 	a.balance += amount
// }

func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balace of your account
func (a Account) Balance() int {
	return a.balance
}

// WithDraw x amount from your account
func (a *Account) WithDraw(amount int) error {
	if a.balance < amount {
		// error 생성
		// return errors.New("돈부족, 불가능")
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "의 계좌 \n남은 잔고: ", a.Balance())
}
