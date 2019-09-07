package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
)

func (spoc *Spoc) devices(endpoint string) {
	b, err := spoc.get(endpoint, nil)
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
		if !flagOnlyIDs {
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

func (spoc *Spoc) play(endpoint, device_id string) {
	params := url.Values{}
	if device_id != "" {
		params.Add("device_id", device_id)
	}
	b, err := spoc.put(endpoint, params, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	fmt.Printf("response: %v\n", string(b))

}
