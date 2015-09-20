package gomuse

import (
	"net/url"
)

type Mortal struct {
	ArtistRef
	BornAt   YMuseTime `json:"born_at"`
	DiedAt   YMuseTime `json:"died_at"`
	RealName string    `json:"realname"`
}

type ArtistRef struct {
	Id      ID            `json:"id"`
	Name    string        `json:"name"`
	Visuals []ObjectImage `json:"visuals"`
}

func (self *ArtistRef) GetURL() *url.URL {
	return getArtistURL(self.Id)
}

type Artist struct {
	ArtistRef
	Mortal
	Bio               string       `json:"bio"`
	Albums            []AlbumBrief `json:"albums"`
	ContributedAlbums []AlbumBrief `json:"contributed_albums"`
}

type ContributionRef struct {
	Name string
	Id   ID
}

type TrackContribution struct {
	ContributionType ContributionRef `json:"contribution_type"`
	Count            uint32          `json:"count"`
}
