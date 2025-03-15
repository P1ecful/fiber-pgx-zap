package fiber

import (
	"context"
	httpmodel "efmo-test/internal/controller/models"
	"efmo-test/internal/models/dto"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"time"
)

func (ctrl *Controller) GetSongLibraryHandler(ctx *fiber.Ctx) error {
	var query httpmodel.GetSongLibraryQueryRequest

	if err := ctx.QueryParser(&query); err != nil {
		ctrl.logger.Debug("can`t to parse query request",
			zap.Any("query", query),
			zap.Error(err),
		)

		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	result, err := ctrl.srv.GetSongLibrary(context.Background(),
		query.OrderByGroup,
		query.OrderByRelease)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(result)
}

func (ctrl *Controller) GetSongTextHandler(ctx *fiber.Ctx) error {
	var query httpmodel.GetTextQueryRequest

	if err := ctx.QueryParser(&query); err != nil {
		ctrl.logger.Debug("can`t to parse query request",
			zap.Any("query", query),
			zap.Error(err),
		)

		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	text, err := ctrl.srv.GetSongText(context.Background(), query.Song, query.Group, query.Verse)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(httpmodel.GetSongTextResponse{
		Text: text,
	})
}

func (ctrl *Controller) DeleteSongHandler(ctx *fiber.Ctx) error {
	var query httpmodel.SongQueryRequest

	if err := ctx.QueryParser(&query); err != nil {
		ctrl.logger.Debug("can`t to parse body request",
			zap.Error(err),
			zap.Any("query", query),
		)

		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	if query.Group == "" || query.Song == "" {
		ctrl.logger.Debug("nil query params in delete handler")
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	if err := ctrl.srv.DeleteSong(context.Background(), query.Song, query.Group); err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (ctrl *Controller) UpdateSongHandler(ctx *fiber.Ctx) error {
	var req httpmodel.UpdateSongRequest
	var query httpmodel.SongQueryRequest

	if err := ctx.QueryParser(&query); err != nil {
		ctrl.logger.Debug("can`t to parse query requests", zap.Error(err))
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	if err := ctx.BodyParser(&req); err != nil {
		ctrl.logger.Debug("can`t to parse body requests", zap.Error(err))
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	formatDate, err := time.Parse("02.01.2006", *req.ReleaseDate)
	if err != nil {
		ctrl.logger.Debug("can`t to parse date body request", zap.Error(err))
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	if err := ctrl.srv.UpdateSong(context.Background(), dto.Song{
		Song:        query.Song,
		Group:       query.Group,
		ReleaseDate: &formatDate,
		Text:        req.Text,
		Link:        req.Link,
	}); err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (ctrl *Controller) AddSongHandler(ctx *fiber.Ctx) error {
	var req httpmodel.AddSongRequest

	if err := ctx.BodyParser(&req); err != nil {
		ctrl.logger.Debug("can`t to parse body requests", zap.Error(err))
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	if err := ctrl.srv.AddSong(context.Background(), req.Song, req.Group); err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (ctrl *Controller) GetSongInfoHandler(ctx *fiber.Ctx) error {
	var query httpmodel.SongQueryRequest

	if err := ctx.QueryParser(&query); err != nil {
		ctrl.logger.Debug("can`t to parse query requests", zap.Error(err))
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	song, err := ctrl.srv.GetSongInfo(context.Background(), query.Song, query.Group)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(httpmodel.GetSongInfoResponse{
		Song:        song.Song,
		Group:       song.Group,
		ReleaseDate: song.ReleaseDate,
		Text:        song.Text,
		Link:        song.Link,
	})
}
