package gomuse

import (
	"net/url"
)

type ImageSize string
type ContentType string
type ResourceImage map[ImageSize]string
type ID string

const (
	Original  ImageSize = "original"
	Thumbnail ImageSize = "thumb"
	Small     ImageSize = "small"
	Medium    ImageSize = "medium"
)

const (
	ArtistCT ContentType = "Artist"
	AlbumCT  ContentType = "Album"
)

type ObjectImage struct {
	Image ResourceImage `json:"image"`
}

type Featureable struct {
	FeaturableId    ID            `json:"featurable_id"`
	FeatureableType ContentType   `json:"featurable_type"`
	Id              ID            `json:"id"`
	Subtitle        string        `json:"subtitle"`
	Title           string        `json:"title"`
	Image           ResourceImage `json:"image"`
}

type Searchable struct {
	SearchableId   ID            `json:"searchable_id"`
	SearchableType ContentType   `json:"searchable_type"`
	SpotifyURI     string        `json:"spotify_uri"`
	Title          string        `json:"title"`
	Album          AlbumBrief    `json:"album"`
	Artist         ArtistRef     `json:"artist"`
	Visuals        []ObjectImage `json:"visuals"`
}

func (self *Featureable) GetURL() *url.URL {
	switch {
	case self.FeatureableType == ArtistCT:
		return getArtistURL(self.FeaturableId)
	case self.FeatureableType == AlbumCT:
		return getAlbumURL(self.FeaturableId)
	}
	return nil
}
