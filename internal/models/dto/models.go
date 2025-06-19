package dto

import "time"

type Song struct {
	SongId      int
	AlbumId     int
	AuthorId    int
	Title       string
	ReleaseDate time.Time
	SongText    *string
	SongUrl     *string
}
