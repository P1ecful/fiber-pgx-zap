package service

import (
	"context"
	"efmo-test/internal/models/dto"
	"efmo-test/internal/service/common"
	"efmo-test/internal/storage"
	"go.uber.org/zap"
	"time"
)

type SongService interface {
	GetSongLibrary(ctx context.Context) ([]dto.Song, error)
	GetSong(ctx context.Context, id int) (dto.Song, error)
	GetSongText(ctx context.Context, id int, verse int) (string, error)
	CreateSong(ctx context.Context, author int, album int, title string, songtext *string, url *string) error
	UpdateSong(ctx context.Context, id int, title *string, text *string, url *string) error
	DeleteSong(ctx context.Context, id int) error
}

type Service struct {
	logger  *zap.Logger
	storage storage.Storage
}

func NewService(logger *zap.Logger, storage storage.Storage) *Service {
	return &Service{
		logger:  logger,
		storage: storage,
	}
}

func (s *Service) GetSongLibrary(ctx context.Context) ([]dto.Song, error) {
	return s.storage.GetSongLibrary(ctx)
}

func (s *Service) GetSong(ctx context.Context, id int) (dto.Song, error) {
	return s.storage.GetSong(ctx, id)
}

func (s *Service) GetSongText(ctx context.Context, id int, verse int) (string, error) {
	fullSongText, err := s.storage.GetSongText(ctx, id)
	if err != nil {
		return "", err
	}

	getVerse, err := common.GetVerse(fullSongText, verse)
	if err != nil {
		s.logger.Debug("failed to get verse from text", zap.Int("id", id), zap.Error(err))
	}

	return getVerse, err
}

func (s *Service) CreateSong(ctx context.Context, author int, album int, title string, songtext *string, url *string) error {
	return s.storage.CreateSong(ctx,
		dto.Song{
			AlbumId:     album,
			AuthorId:    author,
			Title:       title,
			ReleaseDate: time.Now(),
			SongText:    songtext,
			SongUrl:     url,
		})
}

func (s *Service) UpdateSong(ctx context.Context, id int, title *string, text *string, url *string) error {
	return s.storage.UpdateSong(ctx, id, title, text, url)
}

func (s *Service) DeleteSong(ctx context.Context, id int) error {
	return s.storage.DeleteSong(ctx, id)
}
