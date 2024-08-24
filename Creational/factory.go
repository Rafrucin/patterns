package patterns

//FACTORY

import ()

type cartype int

const (
	sedan  cartype = 1
	estate cartype = 2
)

type Car interface {
	printName()
}

type SedanCar struct {
	name string
}

func (sc SedanCar) printName() { println(sc.name) }

type EstateCar struct {
	name string
}

func (ec EstateCar) printName() { println(ec.name) }

func Carfactory(ct cartype, name string) Car {
	switch ct {
	case sedan:
		return SedanCar{name: name}
	case estate:
		return EstateCar{name: name}
	default:
		return nil
	}
}

func Factory() {
	rr := Carfactory(sedan, "125")
	rr.printName()
}