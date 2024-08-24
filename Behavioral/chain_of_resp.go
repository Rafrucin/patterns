package patterns
//Chain Of Responsibility

type step interface {
	execute()
	setNext(step)
}

type opneFlap struct {
	next step
}
func (r *opneFlap) execute() {
	println("flap opened")
	r.next.execute()
}
func (r *opneFlap) setNext(next step) {
	r.next = next
}

type pressStartButton struct {
	next step
}
func (r *pressStartButton) execute() {
	println("button pressed")
	r.next.execute()
}
func (r *pressStartButton) setNext(next step) {
	r.next = next
}

type enterPassword struct {
	next step
}
func (r *enterPassword) execute() {
	println("password entered")
}
func (r *enterPassword) setNext(next step) {
	r.next = next
}


func Chain_of_resp() {
	enterPassword := &enterPassword{}

	pressStartButton := &pressStartButton{}
	pressStartButton.setNext(enterPassword)

	opneFlap := &opneFlap{}
	opneFlap.setNext(pressStartButton)

	opneFlap.execute()
}
