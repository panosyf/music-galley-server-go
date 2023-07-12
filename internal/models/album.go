package models

import (
	"database/sql"
	"time"
)

type Album struct {
	AlbumId  int
	ArtistId int
	Title    string
	Genre    string
	Released time.Time
	Created  time.Time
	Expires  time.Time
}

type AlbumModel struct {
	DB *sql.DB
}

func (m *AlbumModel) InsertAlbum(artistId int, title string, genre string, released string, expires int) (int, error) {
	stmt := `INSERT INTO albums (artist_id, title, genre, released, created, expires)
	VALUES(?, ?, ?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, artistId, title, genre, released, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return (int)(id), nil
}

func (m *AlbumModel) GetAlbum(id int) (*Album, error) {
	return nil, nil
}

func (m *AlbumModel) LatestAlbums() ([]*AlbumModel, error) {
	return nil, nil
}
