package main

type Attacker interface {
	Attack()
}

func main() {
	var att Attacker
	att.Attack() // ERROR
	// build 문법적으로는 문제가 없지만 runtime에러가 생김
	// invalid memory address or nil pointer dereference
}
