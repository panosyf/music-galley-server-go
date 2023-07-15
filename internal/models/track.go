package models

import (
	"database/sql"
	"errors"
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

func (m *TrackModel) GetTrack(trackId int) (*Track, error) {
	stmt := `SELECT artist_id, album_id, title, genre, duration, created, expires FROM tracks
	WHERE expires > UTC_TIMESTAMP() AND artist_id = ?`

	row := m.DB.QueryRow(stmt, trackId)

	t := &Track{}

	err := row.Scan(&t.ArtistId, &t.AlbumId, &t.Title, &t.Genre, &t.Duration, &t.Created, &t.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return t, err
}

func (m *TrackModel) LatestTracks() ([]*TrackModel, error) {
	return nil, nil
}
