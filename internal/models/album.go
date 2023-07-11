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

func (m *AlbumModel) InsertAlbum(title string, genre string, formation int, expires int) (int, error) {
	return 0, nil
}

func (m *AlbumModel) GetAlbum(id int) (*Album, error) {
	return nil, nil
}

func (m *AlbumModel) LatestAlbums() ([]*AlbumModel, error) {
	return nil, nil
}
