package proto

var States struct {
	NeedPhone    Constant
	NeedLocation Constant
	Welcome      Constant
	WaitOrder    Constant
	InWork       Constant
	Offline      Constant
}

func initStates() {
	States.Welcome = "welcome"
	States.NeedPhone = "need-phone"
	States.NeedLocation = "need-location"
	States.WaitOrder = "wait-order"
	States.InWork = "in-work"
	States.Offline = "offline"
}
