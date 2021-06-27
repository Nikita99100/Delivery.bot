package proto

func GetWelcomeButtons() ButtonsSet {
	return ButtonsSet{
		DisplayLocation: Buttons.Display.Inline,
		Buttons: []MsgKeyboardRows{
			{
				MsgButtons: []MsgButton{
					{
						Text: Buttons.Actions.StartWork.S(),
						Data: Buttons.Actions.StartWork.S(),
						Type: Buttons.Type.Regular,
					},
				},
			},
		},
	}
}
func GetNeedPhoneButtons() ButtonsSet {
	return ButtonsSet{
		DisplayLocation: Buttons.Display.Reply,
		Buttons: []MsgKeyboardRows{
			{
				MsgButtons: []MsgButton{
					{
						Text: Buttons.Actions.SendContact.S(),
						Type: Buttons.Type.Contact,
					},
				},
			},
		},
	}
}
func GetFinishWorkButtons() ButtonsSet {
	return ButtonsSet{
		DisplayLocation: Buttons.Display.Reply,
		Buttons: []MsgKeyboardRows{
			{
				MsgButtons: []MsgButton{
					{
						Text: Buttons.Actions.FinishWork.S(),
						Type: Buttons.Type.Regular,
					},
				},
			},
		},
	}
}
func GetErrorButtons() ButtonsSet {
	return ButtonsSet{
		DisplayLocation: Buttons.Display.Inline,
		Buttons: []MsgKeyboardRows{
			{
				MsgButtons: []MsgButton{
					{
						Text: Buttons.Actions.Help.S(),
						Data: Buttons.Actions.Help.S(),
					},
				},
			},
		},
	}
}
func GetNotificationButtons() ButtonsSet {
	return ButtonsSet{
		DisplayLocation: Buttons.Display.Inline,
		Buttons: []MsgKeyboardRows{
			{
				MsgButtons: []MsgButton{
					{
						Text: Buttons.Actions.Accept.S(),
						Data: Buttons.Actions.Accept.S(),
					},
				},
			},
		},
	}
}
func GetTaskButtons(buttonText string) ButtonsSet {
	return ButtonsSet{
		DisplayLocation: Buttons.Display.Inline,
		Buttons: []MsgKeyboardRows{
			{
				MsgButtons: []MsgButton{
					{
						Text: buttonText,
						Data: buttonText,
					},
				},
			},
		},
	}
}
