package handler

import (
	"fmt"
	"github.com/pkg/errors"
	"gitlab.com/faemproject/backend/core/shared/tools"
	"gitlab.com/faemproject/backend/delivery/delivery.bot/models"
	"net/http"
)

func (h *Handler) CheckCourier(chatId string, phoneNum string) (string, error) {
	//return "OK", nil //for debug only
	url := fmt.Sprintf("%s/courier?phone_number=%s&chatid=%s", h.Config.Application.CouriersUrl, phoneNum, chatId)
	status := ""
	err := tools.RPC(http.MethodGet, url, nil, nil, &status)
	return status, errors.Wrap(err, "rpc error")
}

// SetCourierStatusAndLocation кидает put запрос на изменение статуса и координат курьера
func (h *Handler) SetCourierStatusAndLocation(chatId string, status string, coordinates models.Coordinates) error { //hello
	uuid, err := h.GetCourierUUIDByChatId(chatId)
	if err != nil {
		return errors.Wrap(err, "failed to get courier uuid by chatid")
	}
	url := fmt.Sprintf("%s/courier/set_status", h.Config.Application.CouriersUrl)
	courier := models.FillCourier(uuid, status, coordinates)
	err = tools.RPC(http.MethodPut, url, nil, courier, nil)
	if err != nil {
		return errors.Wrap(err, "rpc error")
	}
	return nil
}

func (h *Handler) GetCourierUUIDByChatId(chatId string) (string, error) {
	url := fmt.Sprintf("%s/courier/get_uuid/%s", h.Config.Application.CouriersUrl, chatId)
	uuid := ""
	err := tools.RPC(http.MethodGet, url, nil, nil, &uuid)
	return uuid, errors.Wrap(err, "rpc error")
}

func (h *Handler) NextTask(chatId string) (string, error) {
	uuid, err := h.GetCourierUUIDByChatId(chatId)
	if err != nil {
		return "", errors.Wrap(err, "failed to get courier uuid by chatid")
	}
	url := fmt.Sprintf("%s/next_task/%s", h.Config.Application.CouriersUrl, uuid)
	task := new(models.Task)
	err = tools.RPC(http.MethodGet, url, nil, nil, &task)
	if err != nil {
		return "", errors.Wrap(err, "rpc error")
	}
	address := fmt.Sprint(task.Route.Country, task.Route.City+task.Route.Street+task.Route.House, task.Route.House)
	msg := fmt.Sprintf("⬇️Заказ №%s\n\n📍 %s\n\n📞 %s\n\n💬- %s", task.OrderNumber, address, task.PhoneNumber, task.Route.Comment)
	return msg, nil
}
func (h *Handler) FinishTask(chatId string) error {
	uuid, err := h.GetCourierUUIDByChatId(chatId)
	if err != nil {
		return errors.Wrap(err, "failed to get courier uuid by chatid")
	}
	url := fmt.Sprintf("%s/task/finish/%s", h.Config.Application.CouriersUrl, uuid)
	err = tools.RPC(http.MethodPut, url, nil, nil, nil)
	if err != nil {
		return errors.Wrap(err, "rpc error")
	}
	return nil
}
