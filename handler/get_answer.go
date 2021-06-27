package handler

import (
	"gitlab.com/faemproject/backend/delivery/delivery.bot/models"
	"gitlab.com/faemproject/backend/delivery/delivery.bot/proto"
)

type answerType struct {
	textAnswer    string           //сам текст
	buttonsAnswer proto.ButtonsSet //инлайн кнопки
	msgId         int              //0 - отправить новое, иначе исправить id
}

func (h *Handler) getAnswer(msg *models.ChatMsgFull) (answerType, error) {
	switch msg.State {
	case proto.States.NeedPhone.S():
		return answerType{
			textAnswer:    proto.Consts.BotSend.Answers.SendYourPhoneNumber.S(),
			buttonsAnswer: proto.GetNeedPhoneButtons(),
		}, nil
	case proto.States.Welcome.S():
		return answerType{
			textAnswer:    proto.Consts.BotSend.Answers.Welcome.S(),
			buttonsAnswer: proto.GetWelcomeButtons(),
		}, nil
	case proto.States.NeedLocation.S():
		return answerType{
			textAnswer: proto.Consts.BotSend.Answers.SendYourLocation.S(),
		}, nil
	case proto.States.WaitOrder.S():
		return answerType{
			textAnswer:    proto.Consts.BotSend.Answers.WaitOrder.S(),
			buttonsAnswer: proto.GetFinishWorkButtons(),
		}, nil
	}

	return getErrorAnswer(), nil

}
func getErrorAnswer() answerType {
	return answerType{
		textAnswer:    proto.Consts.BotSend.Answers.ErrorAnswer.S(),
		buttonsAnswer: proto.GetErrorButtons(),
	}
}
func getNextTaskErrorAnswer() answerType {
	return answerType{
		textAnswer: proto.Consts.BotSend.Answers.NextTaskError.S(),
	}
}
func getTaskAnswer(task string) answerType {
	return answerType{
		textAnswer:    task,
		buttonsAnswer: proto.GetTaskButtons(proto.Buttons.Actions.FinishTask.S()),
	}
}
func getCourierNotFoundAnswer() answerType {
	return answerType{
		textAnswer:    proto.Consts.BotSend.Answers.CourierNotFound.S(),
		buttonsAnswer: proto.GetErrorButtons(),
	}
}
