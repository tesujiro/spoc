package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

type PagingAlbums struct {
	PagingBase
	Items []SimplifiedAlbum
}

type SimplifiedAlbum struct {
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
	TotalTracks          int    `json:"total_tracks"`
	Type                 string
	URI                  string
}

type Album struct {
	SimplifiedAlbum
	AlbumGroup  string
	Copyrights  []Copyright
	ExternalIDs ExternalIDs `json:"external_ids"`
	Genres      []string
	Label       string
	Popularity  int
	Tracks      PagingTracks
}

func (album SimplifiedAlbum) String() string {
	if flagOnlyIDs {
		return fmt.Sprintf("%v", album.Id)
	}
	var s string
	s += fmt.Sprintf("%v\t", album.Id)
	s += fmt.Sprintf("name:%v\t", album.Name)
	s += fmt.Sprintf("release:%v\t", album.ReleaseDate)
	s += fmt.Sprintf("tracks:%v\t", album.TotalTracks)
	s += fmt.Sprintf("artists:")
	for _, artist := range album.Artists {
		s += fmt.Sprintf(" %v", artist.Name)
	}
	return s
}

func (album Album) String() string {
	if flagOnlyIDs {
		return fmt.Sprintf("%v", album.Id)
	}
	var s string
	s += fmt.Sprintf("%v\t", album.Id)
	s += fmt.Sprintf("name:%v\t", album.Name)
	s += fmt.Sprintf("release:%v\t", album.ReleaseDate)
	s += fmt.Sprintf("label:%v\t", album.Label)
	s += fmt.Sprintf("artists:")
	for _, artist := range album.Artists {
		s += fmt.Sprintf(" %v", artist.Name)
	}
	s += fmt.Sprintf("\ttracks:")
	for i, track := range album.Tracks.Items {
		s += fmt.Sprintf(" %v:%v", i, track.Name)
	}
	return s
}

func album(token string, endpoint string) {
	b, err := get(token, endpoint, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	var album Album
	err = json.Unmarshal(b, &album)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	fmt.Printf("%v\n", album)
}

func albums(token string, endpoint string, ids []string) {
	maxRec := 20
	for start, end := 0, 0; start < len(ids); start += maxRec {
		if start+maxRec < len(ids) {
			end = start + maxRec
		} else {
			end = len(ids)
		}
		//fmt.Println("start:", start, "end:", end)
		params := url.Values{}
		params.Add("ids", strings.Join(ids[start:end], ","))
		b, err := get(token, endpoint, params)
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
		//fmt.Println(string(b))
		var albums struct {
			Albums []Album
		}
		err = json.Unmarshal(b, &albums)
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
		for i, album := range albums.Albums {
			if flagOnlyIDs {
				fmt.Println(album)
			} else {
				fmt.Printf("Album[%v]: %s\n", start+i, album)
			}
		}
	}
}
