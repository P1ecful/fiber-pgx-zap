package models

import "time"

type GetSongResponse struct {
	SongId      int       `json:"song_id"`
	AlbumId     int       `json:"album_id"`
	AuthorId    int       `json:"author_id"`
	Title       string    `json:"title"`
	ReleaseDate time.Time `json:"release_date"`
	SongText    *string   `json:"song_text"`
	SongUrl     *string   `json:"song_url"`
}
type GetSongTextResponse struct {
	Text string `json:"text"`
}

type AddSongRequest struct {
	AuthorId int     `json:"author_id"`
	AlbumId  int     `json:"album_id"`
	SongText *string `json:"song_text"`
	SongUrl  *string `json:"song_url"`
}

type UpdateSongRequest struct {
	Title    *string `json:"title"`
	SongText *string `json:"song_text"`
	SongUrl  *string `json:"song_url"`
}
