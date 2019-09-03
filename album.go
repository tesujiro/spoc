package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

func album(token string, endpoint string) {
	b, err := get(token, endpoint, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	var album Album
	err = json.Unmarshal(b, &album)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	album.Print()
}

func albums(token string, endpoint string, ids []string) {
	maxRec := 20
	for start, end := 0, 0; start < len(ids); start += maxRec {
		if start+maxRec < len(ids) {
			end = start + maxRec
		} else {
			end = len(ids)
		}
		//fmt.Println("start:", start, "end:", end)
		params := url.Values{}
		params.Add("ids", strings.Join(ids[start:end], ","))
		b, err := get(token, endpoint, params)
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
		//fmt.Println(string(b))
		var albums struct {
			Albums []Album
		}
		err = json.Unmarshal(b, &albums)
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
		for i, album := range albums.Albums {
			fmt.Printf("Album[%v]: ", start+i)
			album.Print()
		}
	}
}
