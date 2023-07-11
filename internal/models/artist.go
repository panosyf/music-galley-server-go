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

func (m *ArtistModel) InsertArtist(name string, genre string, formation int, expires int) (int, error) {
	return 0, nil
}

func (m *ArtistModel) GetArtist(id int) (*Artist, error) {
	return nil, nil
}

func (m *ArtistModel) LatestArtists() ([]*ArtistModel, error) {
	return nil, nil
}
