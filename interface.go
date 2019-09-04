package main

import (
	"time"
)

type Timestamp time.Time

type PagingBase struct {
	Href     string
	Limit    int
	Offset   int
	Total    int
	Next     string
	Previous string
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

type Copyright struct {
	Text string
	Type string
}

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
