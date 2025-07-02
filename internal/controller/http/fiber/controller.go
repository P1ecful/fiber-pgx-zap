package fiber

import (
	"github.com/P1ecful/fiber-pgx-zap/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"go.uber.org/zap"
)

type Controller struct {
	srv    service.SongService
	logger *zap.Logger
	app    *fiber.App
}

func NewController(logger *zap.Logger, srv service.SongService,
	fiber *fiber.App) *Controller {
	return &Controller{
		srv:    srv,
		logger: logger,
		app:    fiber,
	}
}

func (ctrl *Controller) BaseHandler(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusOK)
}

func (ctrl *Controller) ConfigureRoutes() {
	ctrl.app.Get("metrics", monitor.New(monitor.Config{Title: "Fiber-pgx-zap metrics"}))
	ctrl.app.Get("swagger/*", ctrl.BaseHandler)

	song := ctrl.app.Group("song")
	{
		song.Get("library", ctrl.GetSongLibraryHandler)
		song.Get(":song_id", ctrl.GetSongHandler)
		song.Get("text/:song_id", ctrl.GetSongTextHandler)
		song.Post("", ctrl.AddSongHandler)
		song.Put(":song_id", ctrl.UpdateSongHandler)
		song.Delete(":song_id", ctrl.DeleteSongHandler)
	}
}
