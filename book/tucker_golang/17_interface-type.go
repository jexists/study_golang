package main

import "fmt"

type Attacker interface {
	Attack()
}

type Monster struct {
	Lv int
}

func (m *Monster) Attack() {
	fmt.Println("Monster Attack")
}

func DoAttack(att Attacker) {
	if att != nil {
		att.Attack()
		// Monster Attack
		var monster *Monster
		monster = att.(*Monster)
		fmt.Println(att)
		// &{20}
		fmt.Println(monster.Lv)
		// 20
	}
}

func main() {
	DoAttack(&Monster{20})
}
