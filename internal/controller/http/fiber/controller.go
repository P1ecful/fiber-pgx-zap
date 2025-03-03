package fiber

import (
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"srv-tmpl/internal/service"
)

type Controller struct {
	logger *slog.Logger
	srv    *service.Service
	app    *fiber.App
}

func NewController(logger *slog.Logger, srv *service.Service, fiber *fiber.App) *Controller {
	return &Controller{
		logger: logger,
		srv:    srv,
		app:    fiber,
	}
}

func (ctrl *Controller) RegisterRoutes() {
	Methods := ctrl.app.Group("/methods")
	{
		Methods.Get("", ctrl.GetMethodHandler)
		Methods.Post("", ctrl.PostMethodHandler)
	}
}
