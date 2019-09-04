package main

import "fmt"

type PagingTracks struct {
	PagingBase
	Items []Track
}

type TrackLink struct {
	ExternalURLs ExternalURLs `json:"external_urls"`
	Href         string
	Id           string
	Type         string
	URI          string
}

type Track struct {
	Album            Album
	Artists          []Artist
	AvailableMarkets []string `json:"available_markets"`
	DiscNumber       int      `json:"disc_number"`
	DurationMs       int      `json:"duration_ms"`
	Explicit         bool
	ExternalIDs      ExternalIDs  `json:"external_ids"`
	ExternalURLs     ExternalURLs `json:"external_urls"`
	Href             string
	Id               string
	IsPlayable       bool      `json:"is_playable"`
	LinkedFrom       TrackLink `json:"linked_from"`
	Name             string
	Popularity       int
	PreviewURL       string `json:"preview_url"`
	TrackNumber      int    `json:"track_number"`
	Type             string
	URI              string
}

func (track Track) String() string {
	if flagOnlyIDs {
		return fmt.Sprintf("%v\n", track.Id)
	} else {
		var ret string
		ret += fmt.Sprintf("%v\t", track.Id)
		ret += fmt.Sprintf("name:%v\t", track.Name)
		ret += fmt.Sprintf("%v (", track.Name)
		sep := ""
		for _, a := range track.Artists {
			ret += fmt.Sprintf("%v%v", sep, a.Name)
			sep = ", "
		}
		ret += fmt.Sprintf(") album: \"%v\"", track.Album.Name)
		return ret
	}
}
