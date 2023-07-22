package main

import (
	"html/template"
	"path/filepath"

	"github.com/panosyf/music-gallery-server-go/internal/models"
)

type templateData struct {
	Artist  *models.Artist
	Artists []*models.Artist
	Album  *models.Album
	Albums []*models.Album
	Track  *models.Track
	Tracks []*models.Track
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		files := []string{
			"./ui/html/base.tmpl.html",
			"./ui/html/partials/nav.tmpl.html",
			page,
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
