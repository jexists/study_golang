## channel

fixed sized

thread safe queue(FIFO / push, pop)

하나의 타입 (chan)

```go
// 14 channel

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// channel make함수
	c := make(chan int)
	//int타입을 사용하는 채널
	//동기화 작업

	go func() {
		c <- 10
		c <- 20
		c <- 30
	}()

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

	//버퍼 사이즈가 3인 채널 만들기
	bc := make(chan int, 3)
	```go
c := make(chan int, 1)
c<-
```
	bc <- 20
	bc <- 30
	// bc <- 40 //blocking

	fmt.Println(<-bc)
	fmt.Println(<-bc)
	fmt.Println(<-bc)
	// fmt.Println(<-bc) //blocking
	close(bc) //채널 닫기

	// range 연산 가능 (range blocking 될때까지 돌기)
	for val := range bc {
		fmt.Println(val)
	}
}

func Print(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
}
```

### 선언

`var 변수명 chan 변수타입`

```go
var channel chan int
```

### 초기화

`변수명 := make(chan 변수타입, 사이즈)`

```go
//slice 초기화
a := make([]int, 10) 
//map 초기화
a := make(map[int]string) 
//channel
a := make(chan int, 10) //=> fixed size

a := make(chan int) //=> 0개 channel
```

### Push ←

```go
c := make(chan int, 1)
c <- 10
```

### Pop

```go
v := <- c
```

```go
package main

func main() {
	var c chan int        //선언
	c = make(chan int, 1) //초기화

	c <- 10
	v := <-c

	println(v) //10
}
```