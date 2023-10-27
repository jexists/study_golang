package main

var wg sync.WaitGroup
func main() {
	rand.Seed(time.Now().UnixNano())

	wg.Add(2)
	fork := &sync.Mutex{}
	spoon := &sync.Mutex{}

	// go diningProblem("A", fork, spoon, "포크", 수저)
	// go diningProblem("B", spoon, fork, "수저", 포크)
// deadlock
	
	go diningProblem("A",  spoon,fork,  "수저","포크")
	go diningProblem("B", spoon, fork, "수저", "포크")
	// deadlock ok
	wg.Wait()
}

func diningProblem(name string, first, second *sync.Mutex, fisrtName, secondName) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%s 밥을 먹으려 합니다.\n", name)
		first.Lock()
		fmt.Printf("%s %s 획득", name, fisrtName)
		second.Lock()
		fmt.Printf("%s %s 획득", name, secondName)

		fmt.Printf("%s 밥을 먹습니다\n", name)
		time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)

		second.Unlock
		first.*Unlock
	}
	wg.Done()

	// deadlock: 모든 고루틴이 잠들경우
}