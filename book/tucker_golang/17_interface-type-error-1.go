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

type Player struct {
	Lv int
}

func (p *Player) Attack() {
	fmt.Println("Player Attack")
}

func DoAttack(att Attacker) {
	if att != nil {
		att.Attack()
		var monster *Monster
		monster = att.(*Monster)
		fmt.Println(monster.Lv)
		// Runtime Error
		// interface conversion: main.Attacker is *main.Player, not *main.Monster
	}
}

func main() {
	DoAttack(&Player{20})
}
