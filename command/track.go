package command

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/tesujiro/spoc/global"
)

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

type SimplifiedTrack struct {
	Artists          []Artist
	AvailableMarkets []string `json:"available_markets"`
	DiscNumber       int      `json:"disc_number"`
	DurationMs       int      `json:"duration_ms"`
	Explicit         bool
	ExternalURLs     ExternalURLs `json:"external_urls"`
	Href             string
	Id               string
	IsPlayable       bool      `json:"is_playable"`
	LinkedFrom       TrackLink `json:"linked_from"`
	Name             string
	PreviewURL       string `json:"preview_url"`
	TrackNumber      int    `json:"track_number"`
	Type             string
	URI              string
}

type Track struct {
	SimplifiedTrack
	Album       SimplifiedAlbum
	ExternalIDs ExternalIDs `json:"external_ids"`
	Popularity  int
}

//func (track SimplifiedTrack) String() string {
//}

func (track Track) String() string {
	if global.FlagOnlyIDs {
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
		ret += fmt.Sprintf(" release: %v", track.Album.ReleaseDate)
		ret += fmt.Sprintf(" duration: %vsec", track.DurationMs/1000)
		return ret
	}
}

func (cmd *Command) GetTrack(id string) {
	endpoint := strings.ReplaceAll(cmd.endpoint("track"), "{id}", id)
	b, err := cmd.Api.Get(endpoint, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	var track Track
	err = json.Unmarshal(b, &track)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	fmt.Printf("%v\n", track)
}

func (cmd *Command) GetTracks(ids []string) {
	endpoint := cmd.endpoint("tracks")
	maxRec := 50
	for start, end := 0, 0; start < len(ids); start += maxRec {
		if start+maxRec < len(ids) {
			end = start + maxRec
		} else {
			end = len(ids)
		}
		params := url.Values{}
		params.Add("ids", strings.Join(ids[start:end], ","))
		b, err := cmd.Api.Get(endpoint, params)
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
		var tracks struct {
			Tracks []Track
		}
		err = json.Unmarshal(b, &tracks)
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
		for i, track := range tracks.Tracks {
			if global.FlagOnlyIDs {
				fmt.Println(track)
			} else {
				fmt.Printf("Track[%v]: %s\n", start+i, track)
			}
		}
	}
}
