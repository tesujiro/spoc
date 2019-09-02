package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
)

func playlist(token string, endpoint string) {
	b, err := get(token, endpoint, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	//fmt.Printf("%+v\n", string(b))
	var playlist Playlist
	err = json.Unmarshal(b, &playlist)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	// Get all tracks
	tracks := playlist.Tracks.Items
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

	if !flagOnlyIDs {
		//fmt.Printf("Playlist: %+v\n", playlist)
		fmt.Printf("ID: %v\n", playlist.Id)
		fmt.Printf("Desc: %v\n", playlist.Description)
		fmt.Printf("Name: %v\n", playlist.Name)
		//fmt.Printf("Owner: %+v\n", playlist.Owner)
		fmt.Printf("Tracks: %v\n", playlist.Tracks.Total)
	}
	for i, ptrack := range tracks {
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

func playlists(token string, endpoint string) {
	b, err := get(token, endpoint, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	//var playlists PagingPlaylists
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
			fmt.Printf("Playlist[%v]:\t", i)
			fmt.Printf("%v\t", playlist.Id)
			fmt.Printf("tracks:%v\t", playlist.Tracks.Total)
			fmt.Printf("name:%v\t", playlist.Name)
			fmt.Printf("desc:%v\t", playlist.Description)
			fmt.Printf("\n")
		} else {
			fmt.Printf("%v\n", playlist.Id)
		}
	}
}
