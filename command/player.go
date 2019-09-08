package command

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/tesujiro/spoc/global"
)

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
	//fmt.Printf("%#v\n", string(b))
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	for i, device := range ret.Devices {
		if !global.FlagOnlyIDs {
			fmt.Printf("Device[%v]:\t", i)
			fmt.Printf("%v\t", device.Id)
			fmt.Printf("name:%v\t", device.Name)
			fmt.Printf("type:%v\t", device.Type)
			fmt.Printf("vol:%v%%\t", device.VolumePercent)
			fmt.Printf("\n")
		} else {
			fmt.Printf("%v\n", device.Id)
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
