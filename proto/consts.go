package proto

const (
	// ButtonTextSize int = 64 // не проверенно но точно больше 64
	// ButtonDataSize - максимальное количество символов для записи мета данных для кнопки
	ButtonDataSize int = 64
)

// Constant
type Constant string

//S returns string
func (c *Constant) S() string {
	return string(*c)
}

var Consts struct {
	BotSend struct {
		Answers struct {
			SendYourPhoneNumber Constant
			SendYourLocation    Constant
			CourierNotFound     Constant
			ErrorAnswer         Constant
			NextTaskError       Constant
			Welcome             Constant
			WaitOrder           Constant
			FinishWork          Constant
			NewTaskNotification Constant
		}
	}
	ButtonsDisplayLocation struct {
		Inline Constant
		Reply  Constant
	}
	ButtonsTypes struct {
		Regular  Constant
		Contact  Constant
		Location Constant
	}
}

func init() {
	initStates()
	initButtons()
	initStatuses()
	Consts.BotSend.Answers.SendYourPhoneNumber = "Зравствуйте! \nМне нужен ваш номер телефона, чтобы начать"
	Consts.BotSend.Answers.SendYourLocation = "Чтобы начать смену, пришлите livelocation 😝"
	Consts.BotSend.Answers.CourierNotFound = "Вас нет в списке курьеров. Обратитесь в службу тех поддержки"
	Consts.BotSend.Answers.ErrorAnswer = "Что-то пошло не так. Обратитесь в службу тех поддержки"
	Consts.BotSend.Answers.Welcome = "Вы можете начать смену"
	Consts.BotSend.Answers.WaitOrder = "Ищем для вас заказ"
	Consts.BotSend.Answers.NewTaskNotification = "Появился новый заказ!"
	Consts.BotSend.Answers.NextTaskError = "Нет новых заданий"
	Consts.BotSend.Answers.FinishWork = "Новых назначений больше не будет"
}
