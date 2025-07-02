package storage

import (
	"context"

	"github.com/P1ecful/fiber-pgx-zap/internal/models/dto"
)

type Storage interface {
	Ping(ctx context.Context) error
	Disconnect()
	GetSongLibrary(ctx context.Context) ([]dto.Song, error)
	GetSongText(ctx context.Context, id int) (string, error)
	GetSong(ctx context.Context, id int) (dto.Song, error)
	CreateSong(ctx context.Context, song dto.Song) error
	UpdateSong(ctx context.Context, id int, title *string, text *string, url *string) error
	DeleteSong(ctx context.Context, id int) error
}
