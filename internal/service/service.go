package service

import (
	"context"
	"efmo-test/internal/models/dto"
	"efmo-test/internal/service/common"
	"efmo-test/internal/storage"
	"go.uber.org/zap"
)

type EfMoService interface {
	GetSongLibrary(ctx context.Context, group *string, date *string) ([]dto.Song, error)
	GetSongText(ctx context.Context, song string, group string, verse int) (string, error)
	DeleteSong(ctx context.Context, song string, group string) error
	UpdateSong(ctx context.Context, song dto.Song) error
	AddSong(ctx context.Context, song string, group string) error
	GetSongInfo(ctx context.Context, song string, group string) (dto.Song, error)
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

func (s *Service) GetSongLibrary(ctx context.Context, group *string, date *string) ([]dto.Song, error) {
	return s.storage.GetSongList(ctx, group, date)
}

func (s *Service) GetSongText(ctx context.Context, song string, group string, verse int) (string, error) {
	text, err := s.storage.GetSongText(ctx, song, group)
	if err != nil {
		return "", err
	}

	verseFromText, err := common.GetVerse(text, verse)
	if err != nil {
		s.logger.Debug("failed to get verse from text",
			zap.String("song", song),
			zap.Error(err),
		)

		return "", err
	}
	return verseFromText, nil
}

func (s *Service) DeleteSong(ctx context.Context, song string, group string) error {
	return s.storage.DeleteSong(ctx, song, group)
}

func (s *Service) UpdateSong(ctx context.Context, song dto.Song) error {
	return s.storage.UpdateSong(ctx, song)
}

func (s *Service) AddSong(ctx context.Context, song string, group string) error {
	return s.storage.CreateSong(ctx, song, group)
}

func (s *Service) GetSongInfo(ctx context.Context, song string, group string) (dto.Song, error) {
	return s.storage.GetSong(ctx, song, group)
}
