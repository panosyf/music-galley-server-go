package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func CheckPost(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return errors.New("not POST method")
	}
	return nil
}

func CheckId(w http.ResponseWriter, r *http.Request) (int, error) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return id, errors.New("id not applicable")
	}
	return id, nil
}

func homepage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("homepage"))
}

func artist(w http.ResponseWriter, r *http.Request) {
	id, err := CheckId(w, r)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "artist/%d\n", id)
}

func artistCreate(w http.ResponseWriter, r *http.Request) {
	if CheckPost(w, r) != nil {
		return
	}
	fmt.Fprintf(w, "create artist\n")
}

func album(w http.ResponseWriter, r *http.Request) {
	id, err := CheckId(w, r)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "album/%d\n", id)
}

func albumCreate(w http.ResponseWriter, r *http.Request) {
	if CheckPost(w, r) != nil {
		return
	}
	fmt.Fprintf(w, "create album")
}

func track(w http.ResponseWriter, r *http.Request) {
	id, err := CheckId(w, r)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "track/%d\n", id)
}

func trackCreate(w http.ResponseWriter, r *http.Request) {
	if CheckPost(w, r) != nil {
		return
	}
	fmt.Fprintf(w, "create track")
}
