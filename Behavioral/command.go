package patterns

// Command

type command interface {
	execute()
}

type device interface {
	on()
	off()
}

type onCommand struct {
	device device
}

func (oc *onCommand) setDevice(device device) {
	oc.device = device
}
func (oc *onCommand) execute() {
	oc.device.on()
}

type offCommand struct {
	device device
}

func (oc *offCommand) setDevice(device device) {
	oc.device = device
}
func (oc *offCommand) execute() {
	oc.device.off()
}

type tv struct{}

func (t *tv) on() {
	println("tv on")
}
func (t *tv) off() {
	println("tv off")
}

type remote struct {
	cmd command
}

func (r *remote) pressButton() {
	r.cmd.execute()
}
func (r *remote) setCommand(c command) {
	r.cmd = c
}

func Command() {
	tv := &tv{}
	offCommand := &offCommand{}
	onCommand := &onCommand{}

	onCommand.setDevice(tv)
	offCommand.setDevice(tv)

	remote := &remote{}
	remote.setCommand(onCommand)
	remote.pressButton()

	remote.setCommand(offCommand)
	remote.pressButton()

}