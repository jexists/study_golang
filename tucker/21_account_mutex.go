package main

type Account struct {
	Balance int
}

var mutex sync.mutex

func main() {
	var wg sync.WaitGroup

	account := &Account{10}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func DepositAndWithdraw(account *Account) {
	mutex.Lock()
	defer mutex.Unlock()

	if account.Balance <0 {
		panic(fmt.Sprintf("balance should not be negative value %d", account.Balance))
	}

	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}