package main

import "time"

type Timestamp time.Time

type PagingBase struct {
	Href     string
	Limit    int
	Offset   int
	Total    int
	Next     string
	Previous string
}

type PagingAlbums struct {
	PagingBase
	Items []Album
}

type PagingArtists struct {
	PagingBase
	Items []Artist
}

type PagingPlaylists struct {
	PagingBase
	Items []Playlist
}

type PagingPlaylistTracks struct {
	PagingBase
	Items []PlaylistTrack
}

type ExternalIDs map[string]string
type ExternalURLs map[string]string

type Followers struct {
	Href  string
	Total int
}

type Image struct {
	Height int
	URL    string
	Width  int
}

type Restrictions map[string]string

type User struct {
	DisplayName  string
	ExternalURLs ExternalURLs `json:"external_urls"`
	Followers    Followers
	Href         string
	Id           string
	Images       []Image
	Type         string
	URI          string
}

type Device struct {
	Id               string
	IsActive         bool `json:"is_active"`
	IsPrivateSession bool `json:"is_private_session"`
	IsRestricted     bool `json:"is_restricted"`
	Name             string
	Type             string
	VolumePercent    int `json:"volume_percent"`
}

type Album struct {
	AlbumGroup           string
	AlbumType            string
	Artists              []Artist
	AvailableMarkets     []string
	ExternalURLs         ExternalURLs `json:"external_urls"`
	Href                 string
	Id                   string
	Images               []Image
	Name                 string
	ReleaseDate          string `json:"release_date"`
	ReleaseDatePrecision string `json:"release_date_precision"`
	Restrictions         Restrictions
	Type                 string
	URI                  string
}

type Artist struct {
	ExternalURLs ExternalURLs `json:"external_urls"`
	Followers    Followers
	Genres       []string
	Href         string
	Id           string
	Images       []Image
	Name         string
	Popularity   int
	Type         string
	URI          string
}

type Playlist struct {
	Collaborative bool
	Description   string
	ExternalURLs  ExternalURLs `json:"external_urls"`
	Followers     Followers
	Href          string
	Id            string
	Images        []Image
	Name          string
	Owner         User
	Public        bool
	SnapshotId    string `json:"snapshot_id"`
	Tracks        PagingPlaylistTracks
	Type          string
	URI           string
}

type PlaylistTrack struct {
	AddedAt Timestamp
	AddedBy User
	IsLocal bool
	Track   Track
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

type TrackLink struct {
	ExternalURLs ExternalURLs `json:"external_urls"`
	Href         string
	Id           string
	Type         string
	URI          string
}
