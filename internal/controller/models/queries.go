package models

type GetSongLibraryQueryRequest struct {
	OrderByRelease *string `json:"release_date"` //
	OrderByGroup   *string `json:"song_group"`
}

type GetTextQueryRequest struct {
	Song  string `query:"song"`
	Group string `query:"group"`
	Verse int    `query:"verse"`
}

type SongQueryRequest struct {
	Song  string `query:"song"`
	Group string `query:"group"`
}
