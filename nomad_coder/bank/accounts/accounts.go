package accounts

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

// struct making function
// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}
