package patterns

// STATE PATTERN change behaviour based on state 

type tvState interface {
	state()
}

type on struct{}

func (o *on) state() {
	println("TV is ON")
}

type off struct{}
func(o *off) state() {
	println("TV is OFF")
}

type stateContext struct {
	currentTvState tvState
}
func getContex() *stateContext {
	return &stateContext{
		currentTvState: &off{}, //default state
	}
}
func (sc *stateContext) setState(state tvState){
	sc.currentTvState = state
}
func (sc *stateContext) getState(){
	sc.currentTvState.state()
}

func State() {
	tvContex := getContex()
	tvContex.getState()
	tvContex.setState(&on{})
	tvContex.getState()
}
