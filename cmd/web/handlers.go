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

	app.render(w, http.StatusOK, "homepage.tmpl.html", nil)
}

func (app *application) artists(w http.ResponseWriter, r *http.Request) {
	artists, err := app.artistModel.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, http.StatusOK, "artists.tmpl.html", &templateData{
		Artists: artists,
	})
}

func (app *application) artistView(w http.ResponseWriter, r *http.Request) {
	artist_id, err := app.CheckId(w, r)
	if err != nil {
		return
	}

	artist, err := app.artistModel.Get(artist_id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
			return
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.render(w, http.StatusOK, "artistView.tmpl.html", &templateData{
		Artist: artist,
	})
}

func (app *application) artistCreate(w http.ResponseWriter, r *http.Request) {
	if app.CheckPost(w, r) != nil {
		return
	}

	name := "MIW"
	genre := "Metalcore"
	formation := "2004-6-12"
	expires := 365

	id, err := app.artistModel.Insert(name, genre, formation, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/artist/view?id=%d", id), http.StatusSeeOther)
}

func (app *application) albums(w http.ResponseWriter, r *http.Request) {
	albums, err := app.albumModel.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, http.StatusOK, "albums.tmpl.html", &templateData{
		Albums: albums})
}

func (app *application) albumView(w http.ResponseWriter, r *http.Request) {
	albumId, err := app.CheckId(w, r)
	if err != nil {
		return
	}

	album, err := app.albumModel.Get(albumId)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
			return
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.render(w, http.StatusOK, "albumView.tmpl.html", &templateData{
		Album: album})
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

	id, err := app.albumModel.Insert(artistId, title, genre, released, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/album/view?id=%d", id), http.StatusSeeOther)
}

func (app *application) tracks(w http.ResponseWriter, r *http.Request) {
	tracks, err := app.trackModel.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, http.StatusOK, "tracks.tmpl.html", &templateData{
		Tracks: tracks,
	})
}

func (app *application) trackView(w http.ResponseWriter, r *http.Request) {
	trackId, err := app.CheckId(w, r)
	if err != nil {
		return
	}

	track, err := app.trackModel.Get(trackId)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
			return
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.render(w, http.StatusOK, "trackView.tmpl.html", &templateData{
		Track: track,
	})
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

	id, err := app.trackModel.Insert(artistId, albumId, title, genre, duration, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/track/view?id=%d", id), http.StatusSeeOther)
}
