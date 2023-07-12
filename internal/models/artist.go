package models

import (
	"database/sql"
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

func (m *ArtistModel) GetArtist(id int) (*Artist, error) {
	return nil, nil
}

func (m *ArtistModel) LatestArtists() ([]*ArtistModel, error) {
	return nil, nil
}
