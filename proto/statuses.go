package proto

var Statuses struct {
	WaitOrder Constant
	InDeliver Constant
	Offline   Constant
}

func initStatuses() {
	Statuses.WaitOrder = "Ждёт заказ"
	Statuses.InDeliver = "Доставка"
	Statuses.Offline = "Оффлайн"
}
