package fiber

import (
	"context"
	"efmo-test/internal/controller/http/models"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (ctrl *Controller) GetSongLibraryHandler(ctx *fiber.Ctx) error {
	var query models.GetSongLibraryQueryRequest

	if err := ctx.QueryParser(&query); err != nil {
		ctrl.logger.Debug("can`t to parse query request",
			zap.Any("query", query),
			zap.Error(err),
		)

		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	result, err := ctrl.srv.GetSongLibrary(context.Background(), query.Author, query.Album, query.Date)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(result)
}

func (ctrl *Controller) GetSongHandler(ctx *fiber.Ctx) error {
	songId, err := ctx.ParamsInt("song_id")
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	songInfo, err := ctrl.srv.GetSong(context.Background(), songId)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(models.GetSongResponse{
		SongId:      songInfo.SongId,
		AlbumId:     songInfo.AlbumId,
		AuthorId:    songInfo.AuthorId,
		Title:       songInfo.Title,
		ReleaseDate: songInfo.ReleaseDate,
		SongText:    songInfo.SongText,
		SongUrl:     songInfo.SongUrl,
	})
}

func (ctrl *Controller) GetSongTextHandler(ctx *fiber.Ctx) error {
	songId, err := ctx.ParamsInt("song_id")
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	var query models.GetSongTextQueryRequest
	if err := ctx.QueryParser(&query); err != nil {
		ctrl.logger.Debug("can`t to parse query request",
			zap.Any("query", query),
			zap.Error(err),
		)

		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	text, err := ctrl.srv.GetSongText(context.Background(), songId, query.Verse)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(models.GetSongTextResponse{
		Text: text,
	})
}

func (ctrl *Controller) AddSongHandler(ctx *fiber.Ctx) error {
	var req models.AddSongRequest

	if err := ctx.BodyParser(&req); err != nil {
		ctrl.logger.Debug("can`t to parse body requests", zap.Error(err))
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	if err := ctrl.srv.AddSong(
		context.Background(),
		req.AuthorId, req.AlbumId,
		req.SongText,
		req.SongUrl); err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (ctrl *Controller) UpdateSongHandler(ctx *fiber.Ctx) error {
	songId, err := ctx.ParamsInt("song_id")
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	var req models.UpdateSongRequest
	if err := ctx.BodyParser(&req); err != nil {
		ctrl.logger.Debug("can`t to parse body requests", zap.Error(err))
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	if err := ctrl.srv.UpdateSong(
		context.Background(),
		songId,
		req.Title,
		req.SongText,
		req.SongUrl,
	); err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (ctrl *Controller) DeleteSongHandler(ctx *fiber.Ctx) error {
	songId, err := ctx.ParamsInt("song_id")
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	if err := ctrl.srv.DeleteSong(context.Background(), songId); err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.SendStatus(fiber.StatusOK)
}
