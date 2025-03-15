package storage

import (
	"context"
	"efmo-test/internal/models/dto"
)

type Storage interface {
	Ping(ctx context.Context) error
	CreateSong(ctx context.Context, song string, group string) error
	DeleteSong(ctx context.Context, song string, group string) error
	UpdateSong(ctx context.Context, song dto.Song) error
	GetSong(ctx context.Context, song string, group string) (dto.Song, error)
	GetSongList(ctx context.Context, group *string, date *string) ([]dto.Song, error)
	GetSongText(ctx context.Context, song string, group string) (string, error)
	Disconnect()
}
