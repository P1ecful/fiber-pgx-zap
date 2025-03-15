package pgx

import (
	"context"
	"efmo-test/internal/models/dto"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

const (
	createSongQuery = `INSERT INTO song (song_name, song_group) 
						VALUES  (@song, @group);`

	getSongInfoQuery = `SELECT * FROM song 
         				WHERE 
         				    song_group = @group AND 
         				    song_name = @song;`

	getSongTextQuery = `SELECT song_text FROM song 
						WHERE 
						    song_name = @song AND 
	      					song_group = @group;`

	getSongLibraryQuery = `SELECT * FROM song
							WHERE 
		      					($1::text IS NULL OR song_group = $1) AND
		      					($2::date IS NULL OR release_date = $2);`

	deleteSongQuery = `DELETE FROM song 
       						WHERE 
       						    song_name = @song AND 
       						    song_group = @group;`

	updateSongQuery = `UPDATE song 
							SET 
							    	release_date = $1,
									song_text = $2,
									link = $3
							
							WHERE 
							    song_name = $4 AND 
							    song_group = $5;`
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

func (p *PGX) CreateSong(ctx context.Context, song string, group string) error {
	args := pgx.NamedArgs{
		"song":  song,
		"group": group,
	}

	_, err := p.pool.Exec(ctx, createSongQuery, args)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				p.logger.Debug("song exists",
					zap.String("song", song),
					zap.String("group", group),
				)

				return errors.New("song already exists")
			}
		}

		p.logger.Debug("unable to create song", zap.Error(err))
		return err
	}

	return nil
}

func (p *PGX) DeleteSong(ctx context.Context, song string, group string) error {
	if _, err := p.pool.Exec(ctx, deleteSongQuery,
		pgx.NamedArgs{
			"song":  song,
			"group": group,
		}); err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			p.logger.Debug("song not found",
				zap.String("song", song),
				zap.String("group", group),
			)

			return ErrSongNotFound
		}

		p.logger.Debug("unable to delete song", zap.Error(err))
		return err
	}

	return nil
}

func (p *PGX) UpdateSong(ctx context.Context, song dto.Song) error {
	if _, err := p.pool.Exec(ctx, updateSongQuery,
		song.ReleaseDate,
		song.Text,
		song.Link,
		song.Song,
		song.Group); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			p.logger.Debug("song not found", zap.Error(err))
			return ErrSongNotFound
		}

		p.logger.Debug("unable to update song", zap.Error(err))
		return err
	}

	return nil
}

func (p *PGX) GetSong(ctx context.Context, song string, group string) (dto.Song, error) {
	var sng dto.Song
	if err := p.pool.QueryRow(ctx, getSongInfoQuery,
		pgx.NamedArgs{
			"song":  song,
			"group": group,
		}).Scan(
		&sng.Song,
		&sng.Group,
		&sng.ReleaseDate,
		&sng.Text,
		&sng.Link,
	); err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			p.logger.Debug("song not found",
				zap.String("song", song),
				zap.String("group", group),
			)

			return dto.Song{}, ErrSongNotFound
		}

		p.logger.Debug("unable to get song", zap.Error(err))
		return dto.Song{}, err
	}

	return sng, nil
}

func (p *PGX) GetSongList(ctx context.Context, group *string, date *string) ([]dto.Song, error) {
	var songs []dto.Song
	rows, err := p.pool.Query(ctx, getSongLibraryQuery, group, date)
	if err != nil {
		p.logger.Debug("unable to get song list", zap.Error(err))
		return songs, err
	}

	for rows.Next() {
		var song dto.Song

		if err := rows.Scan(
			&song.Song,
			&song.Group,
			&song.ReleaseDate,
			&song.Text,
			&song.Link,
		); err != nil {
			p.logger.Debug("unable to scan song list", zap.Error(err))
			return songs, err
		}

		songs = append(songs, song)
	}

	return songs, nil
}

func (p *PGX) GetSongText(ctx context.Context, song string, group string) (string, error) {
	var text *string

	if err := p.pool.QueryRow(ctx, getSongTextQuery,
		pgx.NamedArgs{
			"song":  song,
			"group": group,
		}).Scan(&text); err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			p.logger.Debug("song not found",
				zap.String("song", song),
				zap.String("group", group),
			)

			return "", ErrSongNotFound
		}

		p.logger.Debug("error getting song text", zap.Error(err))
		return "", err
	}

	if text == nil {
		p.logger.Debug("song text not found",
			zap.String("song", song),
			zap.String("group", group),
		)

		return "", errors.New("song haven`t text")
	}

	return *text, nil
}
