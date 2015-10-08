package gomuse

import (
	"fmt"
	"golang.org/x/oauth2"
	"gopkg.in/jmcvetta/napping.v2"
	"net/http"
	"net/url"
)

type APIError struct{ Code int }

func (self APIError) Error() string {
	return fmt.Sprintf("Got error code %d from API", self.Code)
}

type Client struct {
	ClientId     string
	ClientSecret string
	session      napping.Session
}

type ClientConfig struct {
	ClientId     string
	ClientSecret string
	Debug        bool
}

func NewClient(conf ClientConfig) Client {
	cfg := oauth2.Config{
		ClientID:     conf.ClientId,
		ClientSecret: conf.ClientSecret,
	}
	return Client{
		ClientId:     conf.ClientId,
		ClientSecret: conf.ClientSecret,
		session:      getSession(&cfg, conf.Debug),
	}
}

func (self *Client) GetFeaturedContent() ([]Featureable, error) {
	result := []Featureable{}
	err := APIError{}
	response, errcode := self.session.Get(getFeaturedContentURL().String(), nil, &result, &err)
	e := handleErrors(response, errcode)
	return result, e
}

func (self *Client) GetMortals() ([]Mortal, error) {
	result := []Mortal{}
	err := APIError{}
	response, errcode := self.session.Get(getMortalsURL().String(), nil, &result, &err)
	e := handleErrors(response, errcode)
	return result, e
}

func (self *Client) GetArtist(artistId ID) (Artist, error) {
	result := Artist{}
	err := APIError{}
	response, errcode := self.session.Get(getArtistURL(artistId).String(), nil, &result, &err)
	e := handleErrors(response, errcode)
	return result, e
}

func (self *Client) GetArtistRelations(artistId ID) ([]ArtistRef, error) {
	result := []ArtistRef{}
	err := APIError{}
	response, errcode := self.session.Get(getArtistRelationsURL(artistId).String(), nil, &result, &err)
	e := handleErrors(response, errcode)
	return result, e
}

func (self *Client) GetArtistTrackContributions(artistId ID) ([]TrackContribution, error) {
	result := []TrackContribution{}
	err := APIError{}
	response, errcode := self.session.Get(getArtistTrackContributionsURL(artistId).String(), nil, &result, &err)
	e := handleErrors(response, errcode)
	return result, e
}

func (self *Client) GetAlbum(albumId ID) (Album, error) {
	result := Album{}
	err := APIError{}
	response, errcode := self.session.Get(getAlbumURL(albumId).String(), nil, &result, &err)
	e := handleErrors(response, errcode)
	return result, e
}

func (self *Client) GetTrack(trackId ID) (Track, error) {
	result := Track{}
	err := APIError{}
	response, errcode := self.session.Get(getTrackURL(trackId).String(), nil, &result, &err)
	e := handleErrors(response, errcode)
	return result, e
}

func (self *Client) Search(term string) ([]Searchable, error) {
	result := []Searchable{}
	err := APIError{}
	var params url.Values = make(map[string][]string)
	params.Add("q", term)
	response, errcode := self.session.Get(getSearchURL().String(), &params, &result, err)
	e := handleErrors(response, errcode)
	return result, e
}

func handleErrors(response *napping.Response, errcode error) error {
	if errcode != nil {
		return errcode
	}
	if response.Status() >= 400 {
		return APIError{Code: response.Status()}
	}
	return nil
}

func getSession(cfg *oauth2.Config, debug bool) napping.Session {
	var headers http.Header = make(map[string][]string)
	s := napping.Session{
		Client: getClient(cfg),
		Header: &headers,
		Log:    debug,
	}
	s.Header.Add("Accept", "application/json")
	return s
}
