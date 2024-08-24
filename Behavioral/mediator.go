package patterns

// mediator

type flight interface {
	requestLanding()
	landed()
	permitLanding()
}

type mediator interface {
	canLand(flight) bool
	notify()
}

type flightStruct struct {
	name     string
	mediator mediator
}

func (f *flightStruct) requestLanding() {
	if f.mediator.canLand(f) {
		println(f.name, "is landing")
		f.permitLanding()
	} else {
		println(f.name, "is waiting to land")
	}
}

func (f *flightStruct) landed() {
	println(f.name, "has landed")
	f.mediator.notify()
}

func (f *flightStruct) permitLanding() {
	println(f.name, "got permission to land")
}

type flightControl struct {
	isRunWayFree bool
	landingQueue []flight 
}

func (fc *flightControl) canLand(f flight) bool {
	if fc.isRunWayFree {
		fc.isRunWayFree = false 
		return true
	}
	fc.landingQueue = append(fc.landingQueue, f)
	return false	
}

func (fc *flightControl) notify() {
	if !fc.isRunWayFree {
		fc.isRunWayFree = true
	}
	if len(fc.landingQueue) >0 {
		f := fc.landingQueue[0]
		f.permitLanding()
		fc.landingQueue = fc.landingQueue[1:]
	}
}

func Mediator() {
	fc := &flightControl{
		isRunWayFree: true,
		landingQueue: []flight{},
	}
	f1 := flightStruct{
		name:     "f1",
		mediator: fc,
	}
	f2 := flightStruct{
		name:     "f2",
		mediator: fc,
	}

	f1.requestLanding()
	f2.requestLanding()
	f1.landed()
}
