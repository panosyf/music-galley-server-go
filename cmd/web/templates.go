package main

import "github.com/panosyf/music-gallery-server-go/internal/models"

type artistTemplateData struct {
	Artist  *models.Artist
	Artists []*models.Artist
}

type albumTemplateData struct {
	Album  *models.Album
	Albums []*models.Album
}

type trackTemplateData struct {
	Track  *models.Track
	Tracks []*models.Track
}
