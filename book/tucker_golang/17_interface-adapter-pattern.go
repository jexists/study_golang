package main

import "fmt"

type Database interface {
	Get()
	Set()
	// 아래와 같이 바꾸면 바로 해결되지만 다른곳에 사용하고 있어서  다른 방식으로 해결 -> 어답터 패턴 사용
	// GetData()
	// SetData()
}

// Wrapper Get() Set()을 가지고 있지 않아서 Database 인터페이스 사용불가
type CDatabase struct {
}

func (c CDatabase) GetData() {
	fmt.Println("get")
}

func (c CDatabase) SetData() {
	fmt.Println("set")
}

// Wrapper Get() Set()을 가지고 있어서 Database 인터페이스 사용가능
type Wrapper struct {
	cdb CDatabase
}

func (w Wrapper) Get() {
	w.cdb.GetData()
}

func (w Wrapper) Set() {
	w.cdb.SetData()
}

func main() {
	var cdatabase CDatabase
	var database Database
	database = Wrapper{cdatabase}

	database.Get()

}
