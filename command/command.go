package command

import (
	"fmt"

	"github.com/tesujiro/spoc/api"
)

type Command struct {
	Api *api.Api
}

func New() *Command {
	return &Command{
		Api: api.New(),
	}
}

func Usage() {
	fmt.Println("Usage:")
	fmt.Print(`	spoc search [-id] album(s)|artist(s)|playlist(s)|track(s) [keyword]
	spoc [-id] get album(s) [album_id]+
	spoc [-id] get profile [user_id]+
	spoc [-id] get playlist [playlist_id]+
	spoc [-id] get playlists [user_id]+
	spoc [-id] list device(s)
	spoc [-id] list playlist(s)
	spoc [-id] list profile
	spoc play [device_id]
	spoc pause [device_id]
	spoc playing [device_id]
	spoc play next [device_id]
	spoc play previous [device_id]
`)
}

func (cmd *Command) endpoint(key string) string {
	endpoint := map[string]string{
		"album":         "/v1/albums/{id}",
		"album/tracks":  "/v1/albums/{id}/tracks",
		"albums":        "/v1/albums",
		"devices/me":    "/v1/me/player/devices",
		"search":        "/v1/search",
		"play":          "/v1/me/player",
		"play/me":       "/v1/me/player/play",
		"pause/me":      "/v1/me/player/pause",
		"play/next":     "/v1/me/player/next",
		"play/previous": "/v1/me/player/previous",
		"playlist":      "/v1/playlists/{playlist_id}",
		"playlists/me":  "/v1/me/playlists",
		"playlists":     "/v1/users/{user_id}/playlists",
		"profile/me":    "/v1/me",
		"profile":       "/v1/users/{user_id}",
	}

	return cmd.Api.Base_url + endpoint[key]
}
