package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HomePage)
	mux.HandleFunc("/bands", Bands)
	mux.HandleFunc("/albums", Albums)
	mux.HandleFunc("/songs", Songs)

	log.Println("Starting Server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Homepage"))
}

func Bands(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Band"))
}

func Albums(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Albums"))
}

func Songs(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Songs"))
}
