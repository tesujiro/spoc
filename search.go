package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
)

func (spoc *Spoc) _search(endpoint, target string, args []string) ([]byte, error) {
	params := url.Values{}
	params.Add("limit", fmt.Sprintf("%v", 50))
	params.Add("type", target)
	for _, arg := range args {
		params.Add("q", arg)
	}
	//fmt.Printf("%#v\n", params)
	return spoc.get(endpoint, params)
}

func (spoc *Spoc) search(endpoint string, args []string) {
	if len(args) < 2 {
		usage()
		os.Exit(1)
	}
	switch args[0] {
	case "album", "albums":
		b, err := spoc._search(endpoint, "album", args[1:])
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
		var albums struct {
			Albums PagingAlbums
		}
		err = json.Unmarshal(b, &albums)
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
		if !flagOnlyIDs {
			fmt.Println("Total:", albums.Albums.Total)
		}
		/*
			// get artists info concurrently for each albums
			artists := make([][]Artist, len(albums.Albums.Items))
			wg := sync.WaitGroup{}
			for i, album := range albums.Albums.Items {
				wg.Add(1)
				go func(src []Artist, dst *[]Artist) {
					defer wg.Done()
					for _, artist := range src {
						b, err := get(token, artist.Href, nil)
						if err != nil {
							log.Print(err)
						}
						var a Artist
						err = json.Unmarshal(b, &a)
						if err != nil {
							log.Print(err)
						}
						*dst = append(*dst, a)
					}
				}(album.Artists, &artists[i])
			}
			wg.Wait()
		*/
		// display album info
		for i, album := range albums.Albums.Items {
			if !flagOnlyIDs {
				fmt.Printf("Album[%v]:\t%v\n", i, album)
			} else {
				fmt.Printf("%v\n", album)
			}
		}
	case "artist", "artists":
		b, err := spoc._search(endpoint, "artist", args[1:])
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
		var artists struct {
			Artists PagingArtists
		}
		err = json.Unmarshal(b, &artists)
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
		if !flagOnlyIDs {
			fmt.Println("Total:", artists.Artists.Total)
		}
		for i, artist := range artists.Artists.Items {
			if flagOnlyIDs {
				fmt.Printf("%v\n", artist.Id)
			} else {
				fmt.Printf("Artists[%v]:\t%v\n", i, artist)
			}
		}
	case "playlist", "playlists":
		b, err := spoc._search(endpoint, "playlist", args[1:])
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
		var playlists struct {
			Playlists PagingPlaylists
		}
		err = json.Unmarshal(b, &playlists)
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
		if !flagOnlyIDs {
			fmt.Println("Total:", playlists.Playlists.Total)
		}
		for i, playlist := range playlists.Playlists.Items {
			if !flagOnlyIDs {
				fmt.Printf("Playlist[%v]:\t", i)
				fmt.Printf("%v\t", playlist.Id)
				fmt.Printf("tracks:%v\t", playlist.Tracks.Total)
				fmt.Printf("name:%v", playlist.Name)
				fmt.Printf("\n")
			} else {
				fmt.Printf("%v\n", playlist.Id)
			}
		}
	case "track", "tracks":
		b, err := spoc._search(endpoint, "track", args[1:])
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
		var tracks struct {
			Tracks PagingTracks
		}
		err = json.Unmarshal(b, &tracks)
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
		if !flagOnlyIDs {
			fmt.Println("Total:", tracks.Tracks.Total)
		}
		for i, track := range tracks.Tracks.Items {
			if !flagOnlyIDs {
				fmt.Printf("Track[%v]:\t%v\n", i, track)
			} else {
				fmt.Printf("%v\n", track.Id)
			}
		}
	default:
		usage()
		os.Exit(1)
	}
}
