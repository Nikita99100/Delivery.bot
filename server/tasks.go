package server

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/faemproject/backend/core/shared/logs"
	"net/http"
)

func (r *Rest) NewTaskNotification(c echo.Context) error {
	chatId := c.Param("chat_id")
	err := r.Handler.SendNotification(chatId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, logs.OutputRestError("No such chatId", err))
	}
	return c.JSON(http.StatusOK, logs.OutputRestOK("Notification sended!"))
}
