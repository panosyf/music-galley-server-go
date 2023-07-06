package main

import (
	"fmt"
	"net/http"
)

func (app *application) homepage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/homepage.tmpl.html",
		"./ui/html/partials/nav.tmpl.html"}

	if app.CheckParsedFile(w, files) != nil {
		return
	}
	w.Write([]byte("homepage"))
}

func (app *application) artist(w http.ResponseWriter, r *http.Request) {
	id, err := app.CheckId(w, *r)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "artist/%d\n", id)
}

func (app *application) artistCreate(w http.ResponseWriter, r *http.Request) {
	if app.CheckPost(w, r) != nil {
		return
	}
	fmt.Fprintf(w, "create artist\n")
}

func (app *application) album(w http.ResponseWriter, r *http.Request) {
	id, err := app.CheckId(w, *r)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "album/%d\n", id)
}

func (app *application) albumCreate(w http.ResponseWriter, r *http.Request) {
	if app.CheckPost(w, r) != nil {
		return
	}
	fmt.Fprintf(w, "create album")
}

func (app *application) track(w http.ResponseWriter, r *http.Request) {
	id, err := app.CheckId(w, *r)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "track/%d\n", id)
}

func (app *application) trackCreate(w http.ResponseWriter, r *http.Request) {
	if app.CheckPost(w, r) != nil {
		return
	}
	fmt.Fprintf(w, "create track")
}
