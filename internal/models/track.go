package models

import (
	"database/sql"
	"time"
)

type Track struct {
	TrackId  int
	ArtistId int
	AlbumId  int
	Title    string
	Genre    string
	Duration int
	Created  time.Time
}

type TrackModel struct {
	DB *sql.DB
}

func (m *TrackModel) InsertTrack(title string, genre string, duration int) (int, error) {
	return 0, nil
}

func (m *TrackModel) GetTrack(id int) (*Track, error) {
	return nil, nil
}

func (m *TrackModel) LatestTracks() ([]*TrackModel, error) {
	return nil, nil
}
