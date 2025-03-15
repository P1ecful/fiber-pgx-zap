package fiber

import (
	"efmo-test/internal/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Controller struct {
	srv    service.EfMoService
	logger *zap.Logger
	app    *fiber.App
}

func NewController(logger *zap.Logger, srv service.EfMoService,
	fiber *fiber.App) *Controller {
	return &Controller{
		srv:    srv,
		logger: logger,
		app:    fiber,
	}
}

func (ctrl *Controller) ConfigureRoutes() {
	song := ctrl.app.Group("")
	{
		song.Get("", ctrl.GetSongLibraryHandler)
		song.Get("info", ctrl.GetSongInfoHandler)
		song.Get("text", ctrl.GetSongTextHandler)
		song.Delete("", ctrl.DeleteSongHandler)
		song.Put("", ctrl.UpdateSongHandler)
		song.Post("", ctrl.AddSongHandler)
	}
}
