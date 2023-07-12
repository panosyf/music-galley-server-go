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

	// files := []string{
	// 	"./ui/html/base.tmpl.html",
	// 	"./ui/html/pages/homepage.tmpl.html",
	// 	"./ui/html/partials/nav.tmpl.html"}

	// if app.CheckParsedFile(w, files) != nil {
	// 	return
	// }
	w.Write([]byte("homepage"))
}

func (app *application) artist(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("artist"))
}

func (app *application) artistView(w http.ResponseWriter, r *http.Request) {
	id, err := app.CheckId(w, r)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "artist/view/%d\n", id)
}

func (app *application) artistCreate(w http.ResponseWriter, r *http.Request) {
	if app.CheckPost(w, r) != nil {
		return
	}

	name := "MIW"
	genre := "Metalcore"
	formation := "2004-6-12"
	expires := 365

	id, err := app.artists.InsertArtist(name, genre, formation, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/artist/view?id=%d", id), http.StatusSeeOther)
}

func (app *application) album(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("album"))
}

func (app *application) albumView(w http.ResponseWriter, r *http.Request) {
	id, err := app.CheckId(w, r)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "album/view/%d\n", id)
}

func (app *application) albumCreate(w http.ResponseWriter, r *http.Request) {
	if app.CheckPost(w, r) != nil {
		return
	}

	artistId := 1
	title := "WANYK"
	genre := "Metal"
	released := "2020-9-7"
	expires := 365

	id, err := app.albums.InsertAlbum(artistId, title, genre, released, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/album/view?id=%d", id), http.StatusSeeOther)
}

func (app *application) track(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("track"))
}

func (app *application) trackView(w http.ResponseWriter, r *http.Request) {
	id, err := app.CheckId(w, r)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "track/view/%d\n", id)
}

func (app *application) trackCreate(w http.ResponseWriter, r *http.Request) {
	if app.CheckPost(w, r) != nil {
		return
	}

	artistId := 1
	albumId := 1
	title := "Surfacing"
	genre := "Nu Metal"
	duration := 200
	expires := 365

	id, err := app.tracks.InsertTrack(artistId, albumId, title, genre, duration, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/track/view?id=%d", id), http.StatusSeeOther)
}
