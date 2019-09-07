package main

import (
	"log"
	"os"
	"strings"
)

type Spoc struct {
	token string
}

func NewSpoc() *Spoc {
	//token := ""
	token, err := getAccessToken() // CAUTION: The access tokens expire after 1 hour.
	if err != nil {
		log.Fatal("faild to get access token:", err)
	}
	//fmt.Println("token:", token)

	return &Spoc{token: token}
}

func (spoc *Spoc) Run(cmd string, args []string) {
	//var err error
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
		spoc.search(endpoint["search"], args)
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
				spoc.album(ep)
			default:
				spoc.albums(endpoint["albums"], args)
			}
		case "profile":
			if len(args) == 0 {
				spoc.profile(endpoint["profile/me"])
			} else {
				for _, id := range args {
					ep := strings.ReplaceAll(endpoint["profile"], "{user_id}", id)
					spoc.profile(ep)
				}
			}
		case "playlist":
			for _, id := range args {
				ep := strings.ReplaceAll(endpoint["playlist"], "{playlist_id}", id)
				spoc.playlist(ep)
			}
		case "playlists":
			if len(args) == 0 {
				spoc.playlists(endpoint["playlists/me"])
			} else {
				for _, id := range args {
					ep := strings.ReplaceAll(endpoint["playlists"], "{user_id}", id)
					spoc.playlists(ep)
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
			spoc.devices(endpoint["devices/me"])
		case "playlists", "playlist":
			spoc.playlists(endpoint["playlists/me"])
		case "profile":
			spoc.profile(endpoint["profile/me"])
		default:
			usage()
			os.Exit(1)
		}
	case "play":
		if len(args) == 0 {
			spoc.play(endpoint["play/me"], "")
			return
		}
		switch args[0] {
		case "next":
			if len(args) == 1 {
				spoc.play(endpoint["play/next"], "")
			} else if len(args) == 2 {
				dev := args[1]
				spoc.play(endpoint["play/next"], dev)
			} else {
				usage()
				os.Exit(1)
			}
		case "previous":
			if len(args) == 1 {
				spoc.play(endpoint["play/previous"], "")
			} else if len(args) == 2 {
				dev := args[1]
				spoc.play(endpoint["play/previous"], dev)
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
			spoc.play(endpoint["play/me"], dev)
		}
	default:
		usage()
		os.Exit(1)
	}
}
