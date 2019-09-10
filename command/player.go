package command

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/tesujiro/spoc/global"
)

type CurrentlyPlayingContext struct {
	Device               Device
	Repeat_state         string
	Shuffle_state        bool
	Context              Context
	Timestamp            int
	ProgressMs           int  `json:"progress_ms"`
	IsPlaying            bool `json:"is_playing"`
	Item                 Track
	CurrentlyPlayingType string `json:"currently_playing_type"`
	Actions              struct {
		Disallows struct {
			Interrupting_playback, Pausing, Resuming                         bool
			Seeking, Skipping_next, Skipping_prev                            bool
			Toggling_repeat_context, Toggling_shuffle, Toggling_repeat_track bool
			Transferring_playback                                            bool
		}
	}
}

func (cmd *Command) GetCurrentPlaybackOnDevice(device_id string) {
	endpoint := cmd.endpoint("play")
	params := url.Values{}
	if device_id != "" {
		params.Add("device_id", device_id)
	}
	b, err := cmd.Api.Get(endpoint, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	var ret CurrentlyPlayingContext
	err = json.Unmarshal(b, &ret)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	fmt.Printf("Device:\n\t%v\n", ret.Device)
	fmt.Printf("Track:\n\t%v\n", ret.Item)
}

func (cmd *Command) GetMyDevices() {
	endpoint := cmd.endpoint("devices/me")
	b, err := cmd.Api.Get(endpoint, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	var ret struct {
		Devices []Device
	}
	err = json.Unmarshal(b, &ret)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	for i, device := range ret.Devices {
		if global.FlagOnlyIDs {
			fmt.Println("%v\n", device)
		} else {
			fmt.Printf("Device[%v]:\t%v\n", i, device)
		}
	}
}

func (cmd *Command) play(endpoint, device_id string) {
	params := url.Values{}
	if device_id != "" {
		params.Add("device_id", device_id)
	}
	b, err := cmd.Api.Put(endpoint, params, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	fmt.Printf("response: %v\n", string(b))
}

func (cmd *Command) PlayOnDevice(device_id string) {
	endpoint := cmd.endpoint("play/me")
	cmd.play(endpoint, device_id)
	return
}

func (cmd *Command) PlayNextOnDevice(device_id string) {
	endpoint := cmd.endpoint("play/next")
	cmd.play(endpoint, device_id)
	return
}

func (cmd *Command) PlayPreviousOnDevice(device_id string) {
	endpoint := cmd.endpoint("play/previous")
	cmd.play(endpoint, device_id)
	return
}
