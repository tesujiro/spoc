package command

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/tesujiro/spoc/global"
)

type PlayerError struct {
	Status  int
	Message string
	Reason  string
}

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
	if len(b) == 0 {
		fmt.Println("No available devices are found")
		return
	}
	var ret CurrentlyPlayingContext
	err = json.Unmarshal(b, &ret)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	fmt.Printf("Device:\n\t%v\n", ret.Device)
	fmt.Printf("Track:\n\t%v\n", ret.Item)
	fmt.Printf("Progress:\n\t%vsec\n", ret.ProgressMs/1000)
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

func (cmd *Command) play(endpoint, device_id string, pos_ms int) {
	params := url.Values{}
	if device_id != "" {
		params.Add("device_id", device_id)
	}
	if pos_ms > 0 {
		params.Add("position_ms", fmt.Sprintf("%v", pos_ms))
	}
	b, err := cmd.Api.Put(endpoint, params, nil) // Method: PUT
	if err != nil {
		log.Print(err)
		fmt.Printf("response: %v\n", string(b))
		os.Exit(1)
	}
	//fmt.Printf("response: %v\n", string(b))
}

func (cmd *Command) skip(endpoint, device_id string) {
	params := url.Values{}
	if device_id != "" {
		params.Add("device_id", device_id)
	}
	b, err := cmd.Api.Post(endpoint, params, nil) // Method: POST
	if err != nil {
		log.Print(err)
		fmt.Printf("response: %v\n", string(b))
		os.Exit(1)
	}
	//fmt.Printf("response: %v\n", string(b))
}

func (cmd *Command) PlayOnDevice(device_id string) {
	endpoint := cmd.endpoint("play/me")
	cmd.play(endpoint, device_id, 0)
	return
}

func (cmd *Command) PauseOnDevice(device_id string) {
	endpoint := cmd.endpoint("pause/me")
	cmd.play(endpoint, device_id, 0)
	return
}

func (cmd *Command) SeekOnDevice(device_id string, pos_ms int) {
	//TODO: check the length of currently playing song, if posision is longer than the length, return an error.
	endpoint := cmd.endpoint("seek")
	cmd.play(endpoint, device_id, pos_ms)
	return
}

func (cmd *Command) PlayNextOnDevice(device_id string) {
	endpoint := cmd.endpoint("play/next")
	cmd.skip(endpoint, device_id)
	return
}

func (cmd *Command) PlayPreviousOnDevice(device_id string) {
	endpoint := cmd.endpoint("play/previous")
	cmd.skip(endpoint, device_id)
	return
}
