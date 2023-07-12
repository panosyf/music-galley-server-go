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
	Expires  time.Time
}

type TrackModel struct {
	DB *sql.DB
}

func (m *TrackModel) InsertTrack(artistId int, albumId int, title string, genre string, duration int, expires int) (int, error) {
	stmt := `INSERT INTO tracks (artist_id, album_id, title, genre, duration, created, expires)
	VALUES(?, ?, ?, ?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, artistId, albumId, title, genre, duration, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return (int)(id), nil
}

func (m *TrackModel) GetTrack(id int) (*Track, error) {
	return nil, nil
}

func (m *TrackModel) LatestTracks() ([]*TrackModel, error) {
	return nil, nil
}
