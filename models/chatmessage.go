package models

import (
	"gitlab.com/faemproject/backend/core/shared/structures"
)

const (
	SourceTelegram = "telegram"
)

//ChatMsg global chat message for using inside service
type ChatMsgFull struct {
	structures.MessageFromBot
	Coordinates Coordinates
	State       string `json:"state"`
	Type        string `json:"type"` //тип входящего сообщение
}
