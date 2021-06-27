package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/faemproject/backend/core/shared/jobqueue"
	"gitlab.com/faemproject/backend/delivery/delivery.bot/config"
	"gitlab.com/faemproject/backend/delivery/delivery.bot/proto"
)

const (
	JobQueueNameMsgs  = "income.msg/telegram"
	JobQueueLimitMsgs = 999
)

type TelegramClient interface {
	TelegramInterface
}
type TelegramInterface interface {
	// Отправляем сообщение в чат, опция - клавиатура
	SendMessage(chatID int64, msg string, keyboard ...proto.ButtonsSet) (proto.SendedMessage, error)
	// Обновляем сообщение
	UpdateMessage(chatID int64, msgId int, msg string) (proto.SendedMessage, error)
	// Ставим новую клаву
	UpdateKeyboard(chatID int64, msgID int, newKeyborad proto.ButtonsSet) error
	// Удаляем сообщение целиком
	DeleteMsg(chatID int64, msgID int) error
}
type Handler struct {
	Telegram    TelegramClient
	TelegramBot *tgbotapi.BotAPI
	Jobs        *jobqueue.JobQueues
	Config      *config.Config
}
