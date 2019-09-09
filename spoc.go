package main

import (
	"os"

	"github.com/tesujiro/spoc/command"
)

type Spoc struct {
	Command *command.Command
}

func NewSpoc() *Spoc {
	return &Spoc{
		Command: command.New(),
	}
}

func (spoc *Spoc) Run(cmd string, args []string) {
	switch cmd {
	case "search":
		if len(args) < 2 {
			command.Usage()
			os.Exit(1)
		}
		spoc.Command.Search(args)
	case "get":
		obj := args[0]
		args = args[1:]
		switch obj {
		case "album", "albums":
			switch len(args) {
			case 0:
				command.Usage()
				os.Exit(1)
			case 1:
				id := args[0]
				spoc.Command.GetAlbum(id)
			default:
				ids := args
				spoc.Command.GetAlbums(ids)
			}
		case "profile":
			if len(args) == 0 {
				spoc.Command.GetMyProfile()
			} else {
				for _, id := range args {
					spoc.Command.GetUserProfile(id)
				}
			}
		case "playlist":
			for _, id := range args {
				spoc.Command.GetPlaylist(id)
			}
		case "playlists":
			if len(args) == 0 {
				spoc.Command.GetMyPlaylists()
			} else {
				for _, id := range args {
					spoc.Command.GetUserPlaylists(id)
				}
			}
		default:
			command.Usage()
			os.Exit(1)
		}
	//case "create":
	//obj := args[0]
	//args = args[1:]
	//switch obj {
	//case "playlist":
	//}
	case "list":
		if len(args) > 1 {
			command.Usage()
			os.Exit(1)
		}
		obj := args[0]
		switch obj {
		case "device", "devices":
			spoc.Command.GetMyDevices()
		case "playlists", "playlist":
			spoc.Command.GetMyPlaylists()
		case "profile":
			spoc.Command.GetMyProfile()
		default:
			command.Usage()
			os.Exit(1)
		}
	case "play":
		if len(args) == 0 {
			spoc.Command.PlayOnDevice("")
			return
		}
		switch args[0] {
		case "next":
			if len(args) == 1 {
				spoc.Command.PlayNextOnDevice("")
			} else if len(args) == 2 {
				dev := args[1]
				spoc.Command.PlayNextOnDevice(dev)
			} else {
				command.Usage()
				os.Exit(1)
			}
		case "previous":
			if len(args) == 1 {
				spoc.Command.PlayPreviousOnDevice("")
			} else if len(args) == 2 {
				dev := args[1]
				spoc.Command.PlayPreviousOnDevice(dev)
			} else {
				command.Usage()
				os.Exit(1)
			}
		default:
			if len(args) > 1 {
				command.Usage()
				os.Exit(1)
			}
			dev := args[0]
			spoc.Command.PlayOnDevice(dev)
		}
	default:
		command.Usage()
		os.Exit(1)
	}
}
