package gomuse

type AlbumType string

const (
	LiveAT        AlbumType = "live"
	StudioAT      AlbumType = "studio"
	CompilationAT AlbumType = "compilation"
)

type AlbumBrief struct {
	Id          ID            `json:"id"`
	Artist      ArtistRef     `json:"artist"`
	PublishedOn YMuseTime     `json:"published_on"`
	SpotifyURI  string        `json:"spotify_uri"`
	Title       string        `json:"title"`
	Visuals     []ObjectImage `json:"visuals"`
}

type LabelRef struct {
	Name string
	Id   ID
}

type TrackBrief struct {
	DiscNumber  string `json:"disc_number"`
	Id          ID     `json:"id"`
	SearchTerm  string `json:"search_term"`
	SpotifyURI  string `json:"spotify_uri"`
	Title       string `json:"title"`
	TrackNumber uint16 `json:"track_number"`
}

type Track struct {
	TrackBrief
	Album              AlbumBrief          `json:"album"`
	ContributionGroups []ContributionGroup `json:"contribution_groups"`
	LinerNotes         string              `json:"liner_notes"`
	Lyrics             string              `json:"lyrics"`
}

type ContributionGroup struct {
	Artists []ArtistRef `json:"artists"`
	Name    string      `json:"name"`
}

type Album struct {
	AlbumBrief
	AlbumType          AlbumType           `json:"album_type"`
	ContributionGroups []ContributionGroup `json:"contribution_groups"`
	Label              LabelRef            `json:"label"`
	LinerNotes         string              `json:"liner_notes"`
	LinesNotesHTML     string              `json:"liner_notes_html"`
	Tracks             []TrackBrief        `json:"tracks"`
}
