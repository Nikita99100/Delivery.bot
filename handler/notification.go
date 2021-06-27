package handler

import (
	"gitlab.com/faemproject/backend/delivery/delivery.bot/proto"
	"strconv"
)

func (h *Handler) SendNotification(chatId string) error {
	value, _ := strconv.ParseInt(chatId, 0, 64)
	_, err := h.Telegram.SendMessage(value, proto.Consts.BotSend.Answers.NewTaskNotification.S(), proto.GetNotificationButtons())
	return err
}
