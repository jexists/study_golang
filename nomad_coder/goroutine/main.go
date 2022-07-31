package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [2]string{"joy", "happy"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println("waiting for, ", i)
		fmt.Println(<-c)
	}
}

// func main() {
// 	c := make(chan string)
// 	people := [2]string{"joy", "happy"}
// 	for _, person := range people {
// 		go isSexy(person, c)
// 	}
// 	resultOne := <-c
// 	resultTwo := <-c
// 	// resultThree := <-c
// 	fmt.Println("waiting for message")
// 	fmt.Println("received this message", resultOne)
// 	fmt.Println("received this message", resultTwo)
// 	// fmt.Println("received this message", resultThree) // ERROR channel deallock
// }

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " is sexy"
}

// func main() {
// 	// go sexyCount("joy")
// 	// go sexyCount("happy")

// 	c := make(chan bool)
// 	people := [2]string{"joy", "happy"}
// 	for _, person := range people {
// 		go isSexy(person, c)
// 	}
// 	result := <-c
// 	fmt.Println(result)
// 	fmt.Println(<-c)
// 	// fmt.Println(<-c) ERROR channel deallock
// 	// time.Sleep(time.Second * 10)
// }

// func isSexy(person string, c chan bool) {
// 	time.Sleep(time.Second * 5)
// 	fmt.Println(person)
// 	// true
// 	c <- true
// }

// func isSexy(person string) bool {
// 	time.Sleep(time.Second * 5)
// 	return true
// }

// func sexyCount(person string) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(person, "is sexy", i)
// 		time.Sleep(time.Second)
// 	}
// }
