package patterns

// PROTOTYPE

type ShapeType int

const (
	CircleType ShapeType = 1
	SquareType ShapeType = 2
)

type Shape interface {
	GetType() ShapeType
	PrintTypeProps()
	Clone() Shape
}

type Circle struct {
	Type ShapeType
	Radius int
	Diameter int
	Circumference int
}

func (c Circle) GetType() ShapeType {
	return c.Type
}
func (c Circle) PrintTypeProps() {
	println("printing props", c.Type, c.Circumference)
}
func (c Circle) Clone() Shape {
	return Circle{
		Type:          c.Type,
		Radius:        c.Radius,
		Diameter:      c.Diameter,
		Circumference: c.Circumference,
	}	
}

func Prototype() {
	c := Circle{
		Type:          CircleType,
		Radius:        10,
		Diameter:      1,
		Circumference: 3,
	}

	c2 := c.Clone()

	c.Circumference=66
	c.PrintTypeProps()
	c2.PrintTypeProps()
}