package dto

import "time"

type Song struct {
	Song        string
	Group       string
	ReleaseDate *time.Time
	Text        *string
	Link        *string
}
