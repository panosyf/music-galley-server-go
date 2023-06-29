package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func (app *application) CheckPost(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return errors.New("not POST method")
	}
	return nil
}

func (app *application) CheckId(w http.ResponseWriter, r *http.Request) (int, error) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return id, errors.New("id not applicable")
	}
	return id, nil
}

func (app *application) CheckParsedFile(w http.ResponseWriter, files []string) error {
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return err
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return err
	}
	return nil
}

func (app *application) homepage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
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
	id, err := app.CheckId(w, r)
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
	id, err := app.CheckId(w, r)
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
	id, err := app.CheckId(w, r)
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
