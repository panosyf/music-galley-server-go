package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static")})
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.homepage)

	mux.HandleFunc("/artists", app.artists)
	mux.HandleFunc("/artist/view", app.artistView)
	mux.HandleFunc("/artist/create", app.artistCreate)

	mux.HandleFunc("/albums", app.albums)
	mux.HandleFunc("/album/view", app.albumView)
	mux.HandleFunc("/album/create", app.albumCreate)

	mux.HandleFunc("/tracks", app.tracks)
	mux.HandleFunc("/track/view", app.trackView)
	mux.HandleFunc("/track/create", app.trackCreate)

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(mux)
}

// type neuteredFileSystem struct {
// 	fs http.FileSystem
// }

// func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
// 	f, err := nfs.fs.Open(path)
// 	if err != nil {
// 		return nil, err
// 	}

// 	s, err := f.Stat()
// 	if s.IsDir() {
// 		index := filepath.Join(path, "index.html")
// 		if _, err := nfs.fs.Open(index); err != nil {
// 			if closeErr := f.Close(); closeErr != nil {
// 				return nil, closeErr
// 			}
// 			return nil, err
// 		}
// 	}
// 	return f, nil
// }
