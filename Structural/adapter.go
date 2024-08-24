package patterns

/// Adapter patern - known as WRAPPER two unrelated objects can work together 
type mobile interface {
	chargeAppleMobile()
}

type client struct{}
func (c *client) chargeMobile(mob mobile) {
	mob.chargeAppleMobile()
}
type apple struct{}
func (a *apple) chargeAppleMobile() {
	println("charging APPLE")
}

type android struct{}
func (a *android) chargeAndroidMobile() {
	println("charging ANDROID")
}

type androidAdapter struct {
	android *android
}
func (ad *androidAdapter) chargeAppleMobile() {
	ad.android.chargeAndroidMobile()
}

func Adapter() {

	apple := &apple{}
	client := &client{}
	client.chargeMobile(apple)
	android := &android{}
	//client.chargeMobile(android)
	androidAdapter := &androidAdapter{
		android: android,
	}
	client.chargeMobile(androidAdapter)
}
