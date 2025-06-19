package models

type GetSongLibraryQueryRequest struct {
	Author *int    `query:"author"`
	Album  *int    `query:"album"`
	Date   *string `query:"date"`
}

type GetSongTextQueryRequest struct {
	Verse int `query:"verse"`
}
