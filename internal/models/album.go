package models

import (
	"database/sql"
	"errors"
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

func (m *AlbumModel) GetAlbum(albumId int) (*Album, error) {
	stmt := `SELECT album_id, artist_id, title, genre, released, expires FROM albums
	WHERE expires > UTC_TIMESTAMP() AND album_id = ?`

	row := m.DB.QueryRow(stmt, albumId)

	a := &Album{}

	err := row.Scan(&a.AlbumId, &a.ArtistId, &a.Title, &a.Genre, &a.Released, &a.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return a, nil
}

func (m *AlbumModel) LatestAlbums() ([]*Album, error) {
	stmt := `SELECT album_id, artist_id, title, genre, released, expires FROM albums
	WHERE expires > UTC_TIMESTAMP() ORDER BY album_id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	albums := []*Album{}

	for rows.Next() {
		a := &Album{}
		err := rows.Scan(&a.AlbumId, &a.ArtistId, &a.Title, &a.Genre, &a.Released, &a.Expires)
		if err != nil {
			return nil, err
		}
		albums = append(albums, a)
	}

	if err = rows.Close(); err != nil {
		return nil, err
	}

	return albums, nil
}
