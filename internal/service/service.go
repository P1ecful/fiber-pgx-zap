package service

import (
	"context"
	"efmo-test/internal/models/dto"
	"efmo-test/internal/storage"
	"go.uber.org/zap"
)

type SongService interface {
	GetSongLibrary(ctx context.Context, author *int, album *int, date *string) ([]dto.Song, error)
	GetSong(ctx context.Context, id int) (dto.Song, error)
	GetSongText(ctx context.Context, id int, verse int) (string, error)
	AddSong(ctx context.Context, author int, album int, songtext *string, url *string) error
	UpdateSong(ctx context.Context, id int, title *string, text *string, url *string) error
	DeleteSong(ctx context.Context, song int) error
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

func (s *Service) GetSongLibrary(ctx context.Context, author *int, album *int, date *string) ([]dto.Song, error) {
	return nil, nil
}

func (s *Service) GetSong(ctx context.Context, id int) (dto.Song, error) {
	return dto.Song{}, nil
}

func (s *Service) GetSongText(ctx context.Context, id int, verse int) (string, error) {
	return "", nil
}

func (s *Service) AddSong(ctx context.Context, author int, album int, songtext *string, url *string) error {
	return nil
}

func (s *Service) UpdateSong(ctx context.Context, id int, title *string, text *string, url *string) error {
	return nil
}

func (s *Service) DeleteSong(ctx context.Context, song int) error {
	return nil
}
