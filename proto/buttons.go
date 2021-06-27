package proto

var MsgTypes struct {
	TelegramMessage  Constant
	TelegramCallback Constant
	TelegramContact  Constant
	Coordinates      Constant
	Voice            Constant
}

var Buttons struct {
	Display struct {
		Inline Constant
		Reply  Constant
	}
	Type struct {
		Regular  Constant
		Contact  Constant
		Location Constant
	}

	Actions struct {
		Start       Constant
		SendContact Constant
		FinishWork  Constant
		Help        Constant
		StartWork   Constant
		Accept      Constant
		FinishTask  Constant
	}
}

func initButtons() {
	//Action
	Buttons.Actions.Start = "/start"
	Buttons.Actions.SendContact = "Отправить контакт"
	Buttons.Actions.Help = "Помошь"
	Buttons.Actions.StartWork = "Начать смену"
	Buttons.Actions.Accept = "Принять"
	Buttons.Actions.FinishTask = "Доставил"
	Buttons.Actions.FinishWork = "Завершить смену"
	//Message Types
	MsgTypes.TelegramMessage = "tlgrm_msg"
	MsgTypes.TelegramCallback = "tlgrm_callback"
	MsgTypes.TelegramContact = "tlgrm_contact"
	MsgTypes.Coordinates = "coordinates_msg"
	MsgTypes.Voice = "tlgrm_voice"

	//Расположение кнопок
	Buttons.Display.Inline = "inline"
	Buttons.Display.Reply = "reply"
	//Типы кнопок
	Buttons.Type.Regular = "regular"
	Buttons.Type.Contact = "contact"
	Buttons.Type.Location = "location"
}
