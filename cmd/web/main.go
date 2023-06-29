package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static")})
	mux.Handle("./ui/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.homepage)
	mux.HandleFunc("/artist", app.artist)
	mux.HandleFunc("/artist/create", app.artistCreate)
	mux.HandleFunc("/album", app.album)
	mux.HandleFunc("/album/create", app.albumCreate)
	mux.HandleFunc("/track", app.track)
	mux.HandleFunc("/track/create", app.trackCreate)

	srv := http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
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
