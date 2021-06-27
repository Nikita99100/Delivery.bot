package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/faemproject/backend/delivery/delivery.bot/models"
	"gitlab.com/faemproject/backend/delivery/delivery.bot/proto"
	"strconv"
)

func (h *Handler) sendCoordinates(msg *tgbotapi.Message) error {
	var err error
	if msg.Location != nil {
		chatMsgId := strconv.FormatInt(msg.Chat.ID, 10)
		state := h.GetState(chatMsgId)
		if state == proto.States.WaitOrder.S() || state == proto.States.InWork.S() {
			err = h.SetCourierStatusAndLocation(chatMsgId, "", models.FillCoordinates(
				msg.Location.Latitude, msg.Location.Longitude))
		}
	}
	return err
}
