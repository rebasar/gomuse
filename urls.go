package gomuse

import (
	"fmt"
	"net/url"
)

const (
	apiBase                     string = "https://ymuse.com/api/v1/"
	featuredContentURL          string = "featured_contents"
	mortalsURL                  string = "artists/mortals"
	randomURL                   string = "albums/random?limit=%s"
	searchURL                   string = "search"
	artistURL                   string = "artists/%s"
	artistTrackContributionsURL string = "artists/%s/track_contributions"
	albumURL                    string = "albums/%s"
	trackURL                    string = "tracks/%s"
	oauthTokenURL               string = "https://ymuse.com/oauth/token"
	artistRelationsURL          string = "artists/%s/related"
)

func getFeaturedContentURL() *url.URL {
	result, _ := url.Parse(apiBase + featuredContentURL)
	return result
}

func getMortalsURL() *url.URL {
	result, _ := url.Parse(apiBase + mortalsURL)
	return result
}

func getRandomAlbumsURL(limit uint8) *url.URL {
	formatted := fmt.Sprintf(randomURL, limit)
	result, _ := url.Parse(apiBase + formatted)
	return result
}

func getSearchURL() *url.URL {
	result, _ := url.Parse(apiBase + searchURL)
	return result
}

func getArtistURL(id ID) *url.URL {
	formatted := fmt.Sprintf(artistURL, id)
	result, _ := url.Parse(apiBase + formatted)
	return result
}

func getArtistTrackContributionsURL(id ID) *url.URL {
	formatted := fmt.Sprintf(artistTrackContributionsURL, id)
	result, _ := url.Parse(apiBase + formatted)
	return result
}

func getArtistRelationsURL(id ID) *url.URL {
	formatted := fmt.Sprintf(artistRelationsURL, id)
	result, _ := url.Parse(apiBase + formatted)
	return result
}

func getAlbumURL(id ID) *url.URL {
	formatted := fmt.Sprintf(albumURL, id)
	result, _ := url.Parse(apiBase + formatted)
	return result
}

func getTrackURL(id ID) *url.URL {
	formatted := fmt.Sprintf(trackURL, id)
	result, _ := url.Parse(apiBase + formatted)
	return result
}

func getOauthTokenURL() *url.URL {
	result, _ := url.Parse(oauthTokenURL)
	return result
}
