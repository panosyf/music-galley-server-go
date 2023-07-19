package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/panosyf/music-gallery-server-go/internal/models"
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
	artists, err := app.artists.LatestArtists()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, artist := range artists {
		fmt.Fprintf(w, "%+v\n", artist)
	}
}

func (app *application) artistView(w http.ResponseWriter, r *http.Request) {
	artist_id, err := app.CheckId(w, r)
	if err != nil {
		return
	}

	artist, err := app.artists.GetArtist(artist_id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
			return
		} else {
			app.serverError(w, err)
		}
		return
	}

	fmt.Fprintf(w, "%+v", artist)
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
	albums, err := app.albums.LatestAlbums()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, album := range albums {
		fmt.Fprintf(w, "%+v\n", album)
	}
}

func (app *application) albumView(w http.ResponseWriter, r *http.Request) {
	albumId, err := app.CheckId(w, r)
	if err != nil {
		return
	}

	album, err := app.albums.GetAlbum(albumId)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
			return
		} else {
			app.serverError(w, err)
		}
	}

	fmt.Fprintf(w, "%+v\n", album)
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
	tracks, err := app.tracks.LatestTracks()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, track := range tracks {
		fmt.Fprintf(w, "%+v\n", track)
	}
}

func (app *application) trackView(w http.ResponseWriter, r *http.Request) {
	trackId, err := app.CheckId(w, r)
	if err != nil {
		return
	}

	track, err := app.tracks.GetTrack(trackId)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
			return
		} else {
			app.serverError(w, err)
		}
		return
	}

	fmt.Fprintf(w, "%+v\n", track)
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
