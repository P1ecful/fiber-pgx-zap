package pgx

import (
	"context"
	"efmo-test/internal/models/dto"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

const (
	getSongLibraryQuery = `SELECT * FROM song`
	getSongQuery        = `SELECT * FROM song WHERE id = @id;`
	getSongTextQuery    = `SELECT song_text FROM song WHERE id = @id;`
	createSongQuery     = `INSERT INTO song VALUES  (@id, @album_id, @author_id, @title, @release_date, @song_text, @song_link);`
	updateSongQuery     = `UPDATE song SET title = $1, song_text = $2, song_link = $3 WHERE id = @id;`
	deleteSongQuery     = `DELETE FROM song WHERE id = @id;`
)

var (
	ErrSongNotFound = errors.New("song not found")
)

type PGX struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewPGX(logger *zap.Logger, path string) *PGX {
	pool, err := pgxpool.New(context.Background(), path)

	if err != nil {
		logger.Debug("unable to create connection pool", zap.Error(err))
		return nil
	}

	logger.Info("database initialized and successfully connected")
	return &PGX{
		pool:   pool,
		logger: logger,
	}
}

func (p *PGX) Ping(ctx context.Context) error {
	return p.pool.Ping(ctx)
}

func (p *PGX) Disconnect() {
	p.pool.Close()
}

func (p *PGX) GetSongLibrary(ctx context.Context) ([]dto.Song, error) {
	var songs []dto.Song
	rows, err := p.pool.Query(ctx, getSongLibraryQuery)
	if err != nil {
		p.logger.Debug("unable to get song list", zap.Error(err))
		return songs, err
	}

	for rows.Next() {
		var song dto.Song

		if err := rows.Scan(
			&song.SongId,
			&song.AlbumId,
			&song.AuthorId,
			&song.Title,
			&song.ReleaseDate,
			&song.SongText,
			&song.SongUrl,
		); err != nil {
			p.logger.Debug("unable to scan song list", zap.Error(err))
			return songs, err
		}

		songs = append(songs, song)
	}

	return songs, nil
}

func (p *PGX) GetSongText(ctx context.Context, id int) (string, error) {
	var text *string

	if err := p.pool.QueryRow(ctx, getSongTextQuery, pgx.NamedArgs{"id": id}).Scan(&text); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			p.logger.Debug("song not found", zap.Int("id", id))
			return "", ErrSongNotFound
		}

		p.logger.Debug("error getting song text", zap.Error(err))
		return "", err
	}

	if text == nil {
		p.logger.Debug("no text", zap.Int("id", id))
		return "", errors.New("song haven`t text")
	}

	return *text, nil
}

func (p *PGX) GetSong(ctx context.Context, id int) (dto.Song, error) {
	var sng dto.Song

	if err := p.pool.QueryRow(ctx, getSongQuery, pgx.NamedArgs{"id": id}).Scan(
		&sng.SongId, &sng.AlbumId,
		&sng.AuthorId, &sng.Title,
		&sng.ReleaseDate, &sng.SongText, &sng.SongUrl); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			p.logger.Debug("song not found", zap.Int("id", id))
			return dto.Song{}, ErrSongNotFound
		}

		p.logger.Debug("unable to get song", zap.Error(err))
		return dto.Song{}, err
	}

	return sng, nil
}

func (p *PGX) CreateSong(ctx context.Context, song dto.Song) error {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	args := pgx.NamedArgs{
		"id":           rand.Intn(900000) + 100000,
		"album_id":     song.AlbumId,
		"author_id":    song.AuthorId,
		"title":        song.Title,
		"release_date": song.ReleaseDate.Format("2006-01-02"),
		"song_text":    song.SongText,
		"song_link":    song.SongUrl,
	}

	_, err := p.pool.Exec(ctx, createSongQuery, args)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				p.logger.Debug("song exists")
				return errors.New("song already exists")
			}
		}

		p.logger.Debug("unable to create song", zap.Error(err))
		return err
	}

	return nil
}

func (p *PGX) UpdateSong(ctx context.Context, id int, title *string, text *string, url *string) error {
	args := pgx.NamedArgs{
		"id":    id,
		"title": title,
		"text":  text,
		"url":   url,
	}

	if _, err := p.pool.Exec(ctx, updateSongQuery, args); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			p.logger.Debug("song not found", zap.Error(err))
			return ErrSongNotFound
		}

		p.logger.Debug("unable to update song", zap.Error(err))
		return err
	}

	return nil
}

func (p *PGX) DeleteSong(ctx context.Context, id int) error {
	if _, err := p.pool.Exec(ctx, deleteSongQuery,
		pgx.NamedArgs{"id": id}); err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			p.logger.Debug("song not found", zap.Int("id", id))
			return ErrSongNotFound
		}

		p.logger.Debug("unable to delete song", zap.Error(err))
		return err
	}

	return nil
}
