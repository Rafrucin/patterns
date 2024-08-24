package patterns

//BUILDER

import "fmt"

type house struct {
	wall  string
	roof  string
	doors int
}

type houseBuilder struct {
	wall  string
	roof  string
	doors int
}

func gethouseBuilder() houseBuilder {
	return houseBuilder{}
}
func (b *houseBuilder) buildWalls() {
	b.wall = "done"
}
func (b *houseBuilder) buildRoof() {
	b.roof = "done"
}
func (b *houseBuilder) addDoors() {
	b.doors = 4
}
func (b *houseBuilder) buildHouse() house {
	return house{
		wall:  b.wall,
		roof:  b.roof,
		doors: b.doors,
	}
}

type director struct {
	builder houseBuilder
}
func newDirector(b houseBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) buildHouse() house {
	d.builder.buildWalls()
	d.builder.buildRoof()
	d.builder.addDoors()
	return d.builder.buildHouse()
}

func Builder() {
	houseBuilder := gethouseBuilder()
	director := newDirector(houseBuilder)
	normalHouse := director.buildHouse()
	fmt.Println(normalHouse)
}
