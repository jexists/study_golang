package main

import (
	"fmt"
	"unsafe"
)

type User1 struct {
	Age    int32
	Score  float64
	Age2   int32
	Score2 float64
}
type User2 struct {
	Age    int32
	Age2   int32
	Score  float64
	Score2 float64
}

func main() {
	user1 := User1{}
	user2 := User2{}

	fmt.Println(unsafe.Sizeof(user1)) // 32
	fmt.Println(unsafe.Sizeof(user2)) // 24
}
