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
	Items []Album
}

type Album struct {
	AlbumGroup           string
	AlbumType            string
	Artists              []Artist
	AvailableMarkets     []string
	Copyrights           []Copyright
	ExternalIDs          ExternalIDs  `json:"external_ids"`
	ExternalURLs         ExternalURLs `json:"external_urls"`
	Genres               []string
	Href                 string
	Id                   string
	Images               []Image
	Name                 string
	Popularity           int
	ReleaseDate          string `json:"release_date"`
	ReleaseDatePrecision string `json:"release_date_precision"`
	Tracks               PagingTracks
	Restrictions         Restrictions
	Type                 string
	URI                  string
}

func (album Album) Print() {
	if flagOnlyIDs {
		fmt.Printf("%v\t", album.Id)
		return
	}
	fmt.Printf("%v\t", album.Id)
	fmt.Printf("name:%v\t", album.Name)
	fmt.Printf("release:%v\t", album.ReleaseDate)
	fmt.Printf("artists:")
	for _, artist := range album.Artists {
		fmt.Printf(" %v", artist.Name)
	}
	fmt.Printf("\n")
	fmt.Printf("tracks:")
	for i, track := range album.Tracks.Items {
		fmt.Printf(" %v:%v", i, track.Name)
	}
	fmt.Printf("\n")
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
	album.Print()
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
			fmt.Printf("Album[%v]: ", start+i)
			album.Print()
		}
	}
}
