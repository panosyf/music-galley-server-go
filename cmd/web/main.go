package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", homepage)
	mux.HandleFunc("/artist", artist)
	mux.HandleFunc("/artist/create", artistCreate)
	mux.HandleFunc("/album", album)
	mux.HandleFunc("/album/create", albumCreate)
	mux.HandleFunc("/track", track)
	mux.HandleFunc("/track/create", trackCreate)

	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
