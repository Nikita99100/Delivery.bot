package server

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/faemproject/backend/core/shared/web"
	"gitlab.com/faemproject/backend/delivery/delivery.bot/handler"
)

const (
	apiPrefix = "/api/v1"
)

type Rest struct {
	Router  *echo.Echo
	Handler *handler.Handler
}

// Route defines all the application rest endpoints
func (r *Rest) Route() {
	web.UseHealthCheck(r.Router)
	g := r.Router.Group(apiPrefix)
	g.POST("/newTaskNotification/:chat_id", r.NewTaskNotification)
}
