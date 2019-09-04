package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
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
	var err error
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	f.BoolVar(&flagOnlyIDs, "id", false, "displays only IDs")
	f.BoolVar(&flagRawJson, "json", false, "displays raw json")
	f.Parse(os.Args[1:])
	os.Args = f.Args()

	token, err := getAccessToken() // CAUTION: The access tokens expire after 1 hour.
	if err != nil {
		log.Fatal("faild to get access token:", err)
	}
	//fmt.Println("token:", token)

	cmd := os.Args[0]
	args := os.Args[1:]

	endpoint := map[string]string{
		"album":         base_url + "/v1/albums/{id}",
		"album/tracks":  base_url + "/v1/albums/{id}/tracks",
		"albums":        base_url + "/v1/albums",
		"devices/me":    base_url + "/v1/me/player/devices",
		"search":        base_url + "/v1/search",
		"play/me":       base_url + "/v1/me/player/play",
		"play/next":     base_url + "/v1/me/player/next",
		"play/previous": base_url + "/v1/me/player/previous",
		"playlist":      base_url + "/v1/playlists/{playlist_id}",
		"playlists/me":  base_url + "/v1/me/playlists",
		"playlists":     base_url + "/v1/users/{user_id}/playlists",
		"profile/me":    base_url + "/v1/me",
		"profile":       base_url + "/v1/users/{user_id}",
	}
	switch cmd {
	case "search":
		search(token, endpoint["search"], args)
	case "get":
		obj := args[0]
		args = args[1:]
		switch obj {
		case "album", "albums":
			switch len(args) {
			case 0:
				usage()
				os.Exit(1)
			case 1:
				id := args[0]
				ep := strings.ReplaceAll(endpoint["album"], "{id}", id)
				album(token, ep)
			default:
				albums(token, endpoint["albums"], args)
			}
		case "profile":
			if len(args) == 0 {
				profile(token, endpoint["profile/me"])
			} else {
				for _, id := range args {
					ep := strings.ReplaceAll(endpoint["profile"], "{user_id}", id)
					profile(token, ep)
				}
			}
		case "playlist":
			for _, id := range args {
				ep := strings.ReplaceAll(endpoint["playlist"], "{playlist_id}", id)
				playlist(token, ep)
			}
		case "playlists":
			if len(args) == 0 {
				playlists(token, endpoint["playlists/me"])
			} else {
				for _, id := range args {
					ep := strings.ReplaceAll(endpoint["playlists"], "{user_id}", id)
					playlists(token, ep)
				}
			}
		default:
			usage()
			os.Exit(1)
		}
	case "create":
		obj := args[0]
		args = args[1:]
		switch obj {
		case "playlist":
		}
	case "list":
		if len(args) > 1 {
			usage()
			os.Exit(1)
		}
		obj := args[0]
		switch obj {
		case "device", "devices":
			devices(token, endpoint["devices/me"])
		case "playlists", "playlist":
			playlists(token, endpoint["playlists/me"])
		case "profile":
			profile(token, endpoint["profile/me"])
		default:
			usage()
			os.Exit(1)
		}
	case "play":
		if len(args) == 0 {
			play(token, endpoint["play/me"], "")
			return
		}
		switch args[0] {
		case "next":
			if len(args) == 1 {
				play(token, endpoint["play/next"], "")
			} else if len(args) == 2 {
				dev := args[1]
				play(token, endpoint["play/next"], dev)
			} else {
				usage()
				os.Exit(1)
			}
		case "previous":
			if len(args) == 1 {
				play(token, endpoint["play/previous"], "")
			} else if len(args) == 2 {
				dev := args[1]
				play(token, endpoint["play/previous"], dev)
			} else {
				usage()
				os.Exit(1)
			}
		default:
			if len(args) > 1 {
				usage()
				os.Exit(1)
			}
			dev := args[0]
			play(token, endpoint["play/me"], dev)
		}
	default:
		usage()
		os.Exit(1)
	}
}
