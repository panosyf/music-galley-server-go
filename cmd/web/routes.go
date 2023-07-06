package main

import (
	"net/http"
	"path/filepath"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static")})
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.homepage)

	mux.HandleFunc("/artist", app.artist)
	mux.HandleFunc("/artist/create", app.artistCreate)

	mux.HandleFunc("/album", app.album)
	mux.HandleFunc("/album/create", app.albumCreate)

	mux.HandleFunc("/track", app.track)
	mux.HandleFunc("/track/create", app.trackCreate)

	return mux
}

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			if closeErr := f.Close(); closeErr != nil {
				return nil, closeErr
			}
			return nil, err
		}
	}
	return f, nil
}
