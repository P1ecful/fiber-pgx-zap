package models

import "time"

type AddSongRequest struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

type UpdateSongRequest struct {
	ReleaseDate *string `json:"release_date"`
	Text        *string `json:"text"`
	Link        *string `json:"link"`
}

type GetSongTextResponse struct {
	Text string `json:"text"`
}

type GetSongInfoResponse struct {
	Song        string     `json:"song"`
	Group       string     `json:"group"`
	ReleaseDate *time.Time `json:"release_date"`
	Text        *string    `json:"text"`
	Link        *string    `json:"link"`
}
