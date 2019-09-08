package command

import (
	"fmt"

	"github.com/tesujiro/spoc/global"
)

type PagingArtists struct {
	PagingBase
	Items []Artist
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

func (artist Artist) String() string {
	if global.FlagOnlyIDs {
		return fmt.Sprintf("%v", artist.Id)
	} else {
		var ret string
		ret += fmt.Sprintf("%v\t", artist.Id)
		ret += fmt.Sprintf("name:%v\t", artist.Name)
		return ret
	}
}
