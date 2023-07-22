package models

import (
	"database/sql"
	"errors"
	"time"
)

type Track struct {
	TrackId  int
	ArtistId int
	AlbumId  sql.NullString
	Title    string
	Genre    string
	Duration int
	Created  time.Time
	Expires  time.Time
}

type TrackModel struct {
	DB *sql.DB
}

func (m *TrackModel) Insert(artistId int, albumId int, title string, genre string, duration int, expires int) (int, error) {
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

func (m *TrackModel) Get(trackId int) (*Track, error) {
	stmt := `SELECT track_id, artist_id, album_id, title, genre, duration, created, expires FROM tracks
	WHERE expires > UTC_TIMESTAMP() AND track_id = ?`

	row := m.DB.QueryRow(stmt, trackId)

	t := &Track{}

	err := row.Scan(&t.TrackId, &t.ArtistId, &t.AlbumId, &t.Title, &t.Genre, &t.Duration, &t.Created, &t.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return t, err
}

func (m *TrackModel) Latest() ([]*Track, error) {
	stmt := `SELECT track_id, artist_id, album_id, title, genre, duration, created, expires FROM tracks
	WHERE expires > UTC_TIMESTAMP() ORDER BY track_id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tracks := []*Track{}

	for rows.Next() {
		t := &Track{}
		err := rows.Scan(&t.TrackId, &t.ArtistId, &t.AlbumId, &t.Title, &t.Genre, &t.Duration, &t.Created, &t.Expires)
		if err != nil {
			return nil, err
		}
		tracks = append(tracks, t)
	}

	if err = rows.Close(); err != nil {
		return nil, err
	}

	return tracks, nil
}
