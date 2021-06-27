package handler

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"gitlab.com/faemproject/backend/core/shared/logs"
	"gitlab.com/faemproject/backend/delivery/delivery.bot/models"
	"gitlab.com/faemproject/backend/delivery/delivery.bot/proto"
	"strconv"
	"time"
)

func (h *Handler) TelegramMsgHandler(ctx context.Context, update *tgbotapi.Update) {
	var err error
	log := logs.Eloger.WithFields(logrus.Fields{
		"event": "handling income message ",
	})
	if update.EditedMessage != nil {
		err = h.sendCoordinates(update.EditedMessage)
		if err != nil {
			log.WithFields(logrus.Fields{
				"reason": "stream location error",
			}).Error(err)
		}
		return
	}
	if update.Message == nil && update.CallbackQuery == nil {
		log.WithFields(logrus.Fields{
			"reason": "unknown type. not message and not callback",
		}).Error("failed to response telegram message type")
		return
	}
	msg := fillUpChatMsg(update)
	msg.State = h.GetState(msg.ChatMsgID)
	answer, err := h.handleEventTransition(&msg)
	if err != nil {
		log.WithFields(logrus.Fields{
			"reason": "failed to DFAnswer",
		}).Error(err)
	}
	value, _ := strconv.ParseInt(msg.ChatMsgID, 0, 64)
	_, err = h.Telegram.SendMessage(value, answer.textAnswer, answer.buttonsAnswer)
	if err != nil {
		log.WithFields(logrus.Fields{
			"reason":  "failed to send new msg",
			"chat id": msg.ChatMsgID,
			"msgText": answer.textAnswer,
		}).Error(err)
		return
	}
	h.SetState(msg.ChatMsgID, msg.State)
}

func fillUpChatMsg(update *tgbotapi.Update) models.ChatMsgFull {
	var msg models.ChatMsgFull

	msg.CreatedAt = time.Now()
	msg.Source = models.SourceTelegram

	if update.Message != nil {
		msg.UserLogin = update.Message.From.UserName
		msg.Type = proto.MsgTypes.TelegramMessage.S()
		msg.Text = update.Message.Text
		msg.ChatMsgID = strconv.FormatInt(update.Message.Chat.ID, 10)
		msg.ClientMsgID = strconv.Itoa(update.Message.From.ID)
		msg.CreatedAtMsg = time.Unix(int64(update.Message.Date), 0)
		msg.MsgID = strconv.Itoa(update.Message.MessageID)

		if update.Message.Contact != nil {
			msg.Type = proto.MsgTypes.TelegramContact.S()
			msg.Text = update.Message.Contact.PhoneNumber
			return msg
		}
		if update.Message.Location != nil {
			msg.Type = proto.MsgTypes.Coordinates.S()
			msg.Coordinates.Long = update.Message.Location.Longitude
			msg.Coordinates.Lat = update.Message.Location.Latitude
			return msg
		}
		if update.Message.Voice != nil {
			msg.Type = proto.MsgTypes.Voice.S()
			return msg
		}

		return msg
	}
	//если CallBack
	msg.UserLogin = update.CallbackQuery.From.UserName
	msg.Type = proto.MsgTypes.TelegramCallback.S()
	msg.Text = update.CallbackQuery.Data
	msg.ChatMsgID = strconv.Itoa(int(update.CallbackQuery.Message.Chat.ID))
	msg.ClientMsgID = strconv.Itoa(update.CallbackQuery.From.ID)
	msg.CreatedAtMsg = time.Unix(int64(update.CallbackQuery.Message.Date), 0)
	msg.MsgID = strconv.Itoa(update.CallbackQuery.Message.MessageID)
	return msg
}
