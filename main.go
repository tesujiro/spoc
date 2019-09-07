package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Println("Usage:")
	fmt.Print(`	spoc search [-id] album(s)|artist(s)|playlist(s)|track(s) [keyword]
	spoc [-id] get album(s) [album_id]+
	spoc [-id] get profile [user_id]+
	spoc [-id] get playlist [playlist_id]+
	spoc [-id] get playlists [user_id]+
	spoc [-id] list device(s)
	spoc [-id] list playlist(s)
	spoc [-id] list profile
	spoc [-id] play [device_id]
	spoc [-id] play next [device_id]
	spoc [-id] play previous [device_id]
`)
}

const base_url = "https://api.spotify.com"

var (
	flagOnlyIDs bool
	flagRawJson bool
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	f.BoolVar(&flagOnlyIDs, "id", false, "displays only IDs")
	f.BoolVar(&flagRawJson, "json", false, "displays raw json")
	f.Parse(os.Args[1:])
	os.Args = f.Args()

	spoc := NewSpoc()

	cmd := os.Args[0]
	args := os.Args[1:]

	spoc.Run(cmd, args)
}
