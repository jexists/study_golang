// package main

// func main() {
// 	var wg sync.WaitGroup
// 	ch := make(chan int)

// 	wg.Add(1)
// 	s := []int{}
// 	go square(&wg, s)
// 	s = append(s, 9)

// 	wg.Wait()
// }

// func square(wg *sync.WaitGroup, ch chan int){
// 	n := s[0]

// 	time.Sleep(time.Second)
// 	fmt.Println("Square:", n*n)

// 	wg.Done()
// }
// 문제가 생김

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	s := []int{}
	go square(&wg, s)
	mutex.Lock()
	s = append(s, 9)
	mutex.Unlock()
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	mutex.Lock()
	n := s[0]
	mutex.Unlock()

	time.Sleep(time.Second)
	fmt.Println("Square:", n*n)

	wg.Done()
}
