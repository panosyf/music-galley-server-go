package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HomePage)
	mux.HandleFunc("/artist", Artist)
	mux.HandleFunc("/album", Album)
	mux.HandleFunc("/track", Track)

	mux.HandleFunc("/artist/create", ArtistCreate)
	mux.HandleFunc("/album/create", AlbumCreate)
	mux.HandleFunc("/track/create", TrackCreate)

	log.Println("Starting Server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Homepage"))
}

func Artist(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("/artist"))
}

func CheckPost(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return false
	}
	return true
}

func ArtistCreate(w http.ResponseWriter, r *http.Request) {
	if !CheckPost(w, r) {
		return
	}
	w.Write([]byte("/artist/create"))
}

func Album(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("/album"))
}

func AlbumCreate(w http.ResponseWriter, r *http.Request) {
	if !CheckPost(w, r) {
		return
	}
	w.Write([]byte("/album/create"))
}

func Track(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("/track"))
}

func TrackCreate(w http.ResponseWriter, r *http.Request) {
	if !CheckPost(w, r) {
		return
	}
	w.Write([]byte("/track/create"))
}
