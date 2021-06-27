package handler

import (
	"github.com/pkg/errors"
	"gitlab.com/faemproject/backend/delivery/delivery.bot/models"
	"gitlab.com/faemproject/backend/delivery/delivery.bot/proto"
)

func (h *Handler) handleEventTransition(msg *models.ChatMsgFull) (answerType, error) {
	switch msg.State {
	case proto.States.NeedPhone.S():
		if msg.Type == proto.MsgTypes.TelegramContact.S() {
			status, err := h.CheckCourier(msg.ChatMsgID, msg.Text)
			if err != nil {
				return getErrorAnswer(), errors.Wrap(err, "check courier error")
			}
			if status == "OK" {
				msg.State = proto.States.Welcome.S()
			} else {
				return getCourierNotFoundAnswer(), nil
			}
		}
	case proto.States.Welcome.S():
		if msg.Text == proto.Buttons.Actions.StartWork.S() {
			msg.State = proto.States.NeedLocation.S()
		}
	case proto.States.NeedLocation.S():
		if msg.Type == proto.MsgTypes.Coordinates.S() {
			err := h.SetCourierStatusAndLocation(msg.ChatMsgID, proto.Statuses.WaitOrder.S(), msg.Coordinates)
			if err != nil {
				return getErrorAnswer(),
					errors.Wrap(err, "set courier status error")
			}
			msg.State = proto.States.WaitOrder.S()
		}
	case proto.States.WaitOrder.S():
		if msg.Text == proto.Buttons.Actions.Accept.S() {
			task, err := h.NextTask(msg.ChatMsgID)
			if err != nil {
				return getNextTaskErrorAnswer(), errors.Wrap(err, "failed to get next task")
			}
			msg.State = proto.States.InWork.S()
			return getTaskAnswer(task), nil
		}
		if msg.Text == proto.Buttons.Actions.FinishWork.S() {
			err := h.SetCourierStatusAndLocation(msg.ChatMsgID, proto.Statuses.WaitOrder.S(), models.Coordinates{})
			if err != nil {
				return getErrorAnswer(),
					errors.Wrap(err, "set courier status error")
			}
			msg.State = proto.States.Offline.S()
			return answerType{
				textAnswer:    proto.Consts.BotSend.Answers.FinishWork.S(),
				buttonsAnswer: proto.GetWelcomeButtons(),
			}, nil
		}
	case proto.States.InWork.S():
		if msg.Text == proto.Buttons.Actions.FinishTask.S() {
			err := h.FinishTask(msg.ChatMsgID)
			if err != nil {
				return getErrorAnswer(), errors.Wrap(err, "failed to finish task")
			}
			task, err := h.NextTask(msg.ChatMsgID)
			if err != nil {
				msg.State = proto.States.WaitOrder.S()
				return getNextTaskErrorAnswer(), errors.Wrap(err, "failed to get next task")
			}
			msg.State = proto.States.InWork.S()
			return getTaskAnswer(task), nil
		}

	}

	answer, err := h.getAnswer(msg)
	if err != nil {
		return answerType{}, errors.Wrap(err, "get answer failed")
	}
	return answer, nil
}
