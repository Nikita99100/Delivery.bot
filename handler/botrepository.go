package handler

import "gitlab.com/faemproject/backend/delivery/delivery.bot/proto"

var states map[string]string

func init() {
	states = make(map[string]string)
}
func (h *Handler) GetState(userID string) string {
	val, ok := states[userID]
	if ok {
		return val
	}
	return proto.States.NeedPhone.S()
}
func (h *Handler) SetState(userID string, state string) {
	states[userID] = state
}
