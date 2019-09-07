package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type PagingPlaylists struct {
	PagingBase
	Items []Playlist
}

type PagingPlaylistTracks struct {
	PagingBase
	Items []PlaylistTrack
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

func (playlist Playlist) String() string {
	if flagOnlyIDs {
		return fmt.Sprintf("%v\n", playlist.Id)
	} else {
		var ret string
		ret += fmt.Sprintf("%v\t", playlist.Id)
		ret += fmt.Sprintf("tracks:%v\t", playlist.Tracks.Total)
		ret += fmt.Sprintf("name:%v\t", playlist.Name)
		ret += fmt.Sprintf("desc:%v\t", playlist.Description)
		return ret
	}
}

func (playlist Playlist) PrintDetail() {
	if !flagOnlyIDs {
		fmt.Printf("ID: %v\n", playlist.Id)
		fmt.Printf("Desc: %v\n", playlist.Description)
		fmt.Printf("Name: %v\n", playlist.Name)
		//fmt.Printf("Owner: %+v\n", playlist.Owner)
		fmt.Printf("Tracks: %v\n", playlist.Tracks.Total)
	}
	for i, ptrack := range playlist.Tracks.Items {
		if !flagOnlyIDs {
			fmt.Printf("Track[%v]:\t", i)
			fmt.Printf("%v\t", ptrack.Track.Id)
			fmt.Printf("%v (", ptrack.Track.Name)
			sep := ""
			for _, a := range ptrack.Track.Album.Artists {
				fmt.Printf("%v%v", sep, a.Name)
				sep = ", "
			}
			fmt.Printf(") album: \"%v\"\n", ptrack.Track.Album.Name)
		} else {
			fmt.Printf("%v\n", ptrack.Track.Id)
		}
	}
}

func (spoc *Spoc) playlist(endpoint string) {
	b, err := spoc.get(endpoint, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	var playlist Playlist
	err = json.Unmarshal(b, &playlist)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	/*
		// Get all tracks
		tracks := playlist.Tracks.Items
		//var tracks PagingPlaylistTracks
		total := playlist.Tracks.Total
		limit := playlist.Tracks.Limit
		endpoint += "/tracks" //TODO: not good
		params := url.Values{}
		params.Add("limit", fmt.Sprintf("%v", limit))
		for i := 0; i < total/limit; i++ {
			params.Add("offset", fmt.Sprintf("%v", limit*(i+1)))
			b, err := get(token, endpoint, params)
			if err != nil {
				log.Print(err)
				os.Exit(1)
			}
			var delta PagingPlaylistTracks
			err = json.Unmarshal(b, &delta)
			if err != nil {
				log.Print(err)
				os.Exit(1)
			}
			for _, track := range delta.Items {
				tracks = append(tracks, track)
			}
		}
	*/
	playlist.PrintDetail()
}

func (spoc *Spoc) playlists(endpoint string) {
	b, err := spoc.get(endpoint, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	var playlists struct {
		PagingBase
		Items []Playlist
	}
	err = json.Unmarshal(b, &playlists)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	if !flagOnlyIDs {
		fmt.Println("Total:", playlists.Total)
	}
	for i, playlist := range playlists.Items {
		if !flagOnlyIDs {
			fmt.Printf("Playlist[%v]:\t%v\n", i, playlist)
		} else {
			fmt.Printf("%v\n", playlist)
		}
	}
}
