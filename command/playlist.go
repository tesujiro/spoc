package command

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tesujiro/spoc/global"
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
	Collaborative *bool
	Description   *string
	ExternalURLs  ExternalURLs `json:"external_urls"`
	Followers     Followers
	Href          string
	Id            string
	Images        []Image
	Name          string
	Owner         User
	Public        *bool
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
	if global.FlagOnlyIDs {
		return fmt.Sprintf("%v\n", playlist.Id)
	} else {
		var ret string
		ret += fmt.Sprintf("%v", playlist.Id)
		//if playlist.Public != nil {
		//ret += fmt.Sprintf("public:%v\t", *playlist.Public)
		//}
		//if playlist.Collaborative != nil {
		//ret += fmt.Sprintf("collaborative:%v\t", *playlist.Collaborative)
		//}
		ret += fmt.Sprintf("\ttracks:%v", playlist.Tracks.Total)
		ret += fmt.Sprintf("\tname:%v", playlist.Name)
		if playlist.Description != nil {
			ret += fmt.Sprintf("\tdesc:%v", *playlist.Description)
		}
		return ret
	}
}

func (playlist Playlist) PrintDetail() {
	if !global.FlagOnlyIDs {
		fmt.Printf("ID: %v\n", playlist.Id)
		if playlist.Public != nil {
			fmt.Printf("public:%v\t", *playlist.Public)
		}
		if playlist.Collaborative != nil {
			fmt.Printf("collaborative:%v\t", *playlist.Collaborative)
		}
		if playlist.Description != nil {
			fmt.Printf("Desc: %v\n", *playlist.Description)
		}
		fmt.Printf("Name: %v\n", playlist.Name)
		//fmt.Printf("Owner: %+v\n", playlist.Owner)
		fmt.Printf("Tracks: %v\n", playlist.Tracks.Total)
	}
	for i, ptrack := range playlist.Tracks.Items {
		if !global.FlagOnlyIDs {
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

func (cmd *Command) GetPlaylist(id string) {
	endpoint := strings.ReplaceAll(cmd.endpoint("playlist"), "{playlist_id}", id)
	b, err := cmd.Api.Get(endpoint, nil)
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

func (cmd *Command) getPlaylists(endpoint string) {
	b, err := cmd.Api.Get(endpoint, nil)
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
	if !global.FlagOnlyIDs {
		fmt.Println("Total:", playlists.Total)
	}
	for i, playlist := range playlists.Items {
		if !global.FlagOnlyIDs {
			fmt.Printf("Playlist[%v]:\t%v\n", i, playlist)
		} else {
			fmt.Printf("%v\n", playlist)
		}
	}
}

func (cmd *Command) GetMyPlaylists() {
	endpoint := cmd.endpoint("playlists/me")
	cmd.getPlaylists(endpoint)
}

func (cmd *Command) GetUserPlaylists(id string) {
	endpoint := strings.ReplaceAll(cmd.endpoint("playlists"), "{user_id}", id)
	cmd.getPlaylists(endpoint)
}
