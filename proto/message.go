package proto

type (
	MsgKeyboardRows struct {
		MsgButtons []MsgButton
	}

	MsgButton struct {
		Text string // то что отобразится на кнопке
		Data string
		Type Constant
	}

	MessangerContact struct {
		PhoneNumber string `json:"phone_number"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"` // optional
		UserID      string `json:"user_id"`   // optional
	}

	ButtonsSet struct {
		DisplayLocation Constant // место отображения (в чате или под строкой ввода)
		Buttons         []MsgKeyboardRows
	}

	SendedMessage struct {
		Id string
	}
)

// // ButtonData - для записи в data (параметр кнопки)
// type ButtonData struct {
// 	Action Constant //
// 	Value  string   // значение которое дает кнопка
// }
