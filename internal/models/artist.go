package models

import (
	"database/sql"
	"errors"
	"time"
)

type Artist struct {
	ArtistId  int
	Name      string
	Genre     string
	Formation time.Time
	Created   time.Time
	Expires   time.Time
}

type ArtistModel struct {
	DB *sql.DB
}

func (m *ArtistModel) InsertArtist(name string, genre string, formation string, expires int) (int, error) {
	stmt := `INSERT INTO artists (name, genre, formation, created, expires)
	VALUES(?, ?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, name, genre, formation, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return (int)(id), nil
}

func (m *ArtistModel) GetArtist(artistId int) (*Artist, error) {
	stmt := `SELECT artist_id, name, genre, formation, expires FROM artists
	WHERE expires > UTC_TIMESTAMP() AND artist_id = ?`

	row := m.DB.QueryRow(stmt, artistId)

	a := &Artist{}

	err := row.Scan(&a.ArtistId, &a.Name, &a.Genre, &a.Formation, &a.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return a, nil
}

func (m *ArtistModel) LatestArtists() ([]*Artist, error) {
	stmt := `SELECT artist_id, name, genre, formation, expires FROM artists
	WHERE expires > UTC_TIMESTAMP ORDER BY artist_id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	artists := []*Artist{}

	for rows.Next() {
		a := &Artist{}
		err := rows.Scan(&a.ArtistId, &a.Name, &a.Genre, &a.Formation, &a.Expires)
		if err != nil {
			return nil, err
		}
		artists = append(artists, a)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return artists, nil
}
