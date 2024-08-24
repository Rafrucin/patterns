package patterns

import (
	"fmt"
	"math/rand"
)

//memento
type player struct {
	hp int
	dist_total int
	name string
}

func (p *player) fight() {
	hit := rand.Intn(10)
	if hit >= p.hp {
		println( p.name, "died")
	} else {
		p.hp -= hit
	}
}

func (p *player) walk(dist int) {
	p.dist_total += dist
}

type gameSaver struct{
	saves map[int]player
	round int
}

func newGameSaver() *gameSaver{
	return &gameSaver{
		saves: map[int]player{},
		round: 0,
	}
}

func (sg *gameSaver) restor (round int, p *player) {
	if sg.round >= round {
		println("round not played")
	}
	sg.round = round
	*p = sg.saves[round]
}

func (sg *gameSaver) save(p player) {
	sg.saves[sg.round] = p
	sg.round++
}

func Memento() {
	p1 := player{
		hp:         100,
		dist_total: 0,
		name:       "RR",
	}

	sg := newGameSaver()
	sg.save(p1)
	fmt.Println(p1)
	p1.fight()
	p1.walk(20)
	sg.save(p1)
	fmt.Println(p1)

	p1.fight()
	p1.walk(20)
	sg.save(p1)
	fmt.Println(p1)

	sg.restor(1, &p1)

	fmt.Println("afer restore",p1)
}