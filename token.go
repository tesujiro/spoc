package main

import (
	"crypto/rand"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"

	"golang.org/x/oauth2"
)

func openbrowser(rawurl string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", rawurl).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", rawurl).Start()
	case "darwin":
		err = exec.Command("open", rawurl).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func save(token string, timestamp *time.Time, filename string) error {
	fd, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fd.Close()

	enc := gob.NewEncoder(fd)
	err = enc.Encode(token)
	if err != nil {
		return err
	}
	err = enc.Encode(timestamp)
	if err != nil {
		return err
	}
	return nil
}

func load(filename string) (string, *time.Time, error) {
	fd, err := os.Open(filename)
	if err != nil {
		return "", nil, err
	}
	defer fd.Close()

	dec := gob.NewDecoder(fd)
	var token string
	err = dec.Decode(&token)
	if err != nil {
		return "", nil, err
	}
	var timestamp *time.Time
	err = dec.Decode(&timestamp)
	if err != nil {
		return "", nil, err
	}
	return token, timestamp, nil
}

const validDuration = 1 * time.Hour

func getAccessToken() (string, error) {
	filename := "token.gob"
	token, timestamp, err := load(filename)
	if err == nil {
		if time.Now().Before(timestamp.Add(validDuration)) {
			return token, nil
		}
	} else if err != nil && !os.IsNotExist(err) {
		fmt.Printf("load index error: %v\n", err)
		return "", err
	}
	now := time.Now()

	l, err := net.Listen("tcp", "localhost:8989")
	if err != nil {
		return "", err
	}
	defer l.Close()

	clientID := os.Getenv("ClientID")
	clientSecret := os.Getenv("ClientSecret")
	if clientID == "" || clientSecret == "" {
		err := fmt.Errorf("Env \"ClientID\", \"ClientSecret\" is not set")
		log.Fatal(err)
		return "", err
	}
	conf := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes: []string{
			// CAUTION: set scopes for APIs
			"user-read-playback-state",
			"playlist-read-private",
			"user-modify-playback-state",
		},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.spotify.com/authorize",
			TokenURL: "https://accounts.spotify.com/api/token",
		},
		RedirectURL: "http://localhost:8989", // CAUTION: this URL must be set on the Spotify dashboard
	}

	stateBytes := make([]byte, 16)
	_, err = rand.Read(stateBytes)
	if err != nil {
	}

	state := fmt.Sprintf("%x", stateBytes)
	//rawurl := conf.AuthCodeURL(state, oauth2.AccessTypeOffline)
	rawurl := conf.AuthCodeURL(state, oauth2.SetAuthURLParam("response_type", "token"))
	//fmt.Println("URL:", rawurl)

	// open in browser
	err = openbrowser(rawurl)
	if err != nil {
		return "", err
	}

	// Get Access token with some hacking
	// see https://mattn.kaoriya.net/software/lang/go/20161231001721.htm
	// see https://qiita.com/TakahikoKawasaki/items/8567c80528da43c7e844#%E3%83%95%E3%83%A9%E3%82%B0%E3%83%A1%E3%83%B3%E3%83%88%E9%83%A8%E3%81%AF-http-%E3%83%AA%E3%82%AF%E3%82%A8%E3%82%B9%E3%83%88%E3%81%AB%E5%90%AB%E3%81%BE%E3%82%8C%E3%81%AA%E3%81%84
	quit := make(chan string)
	go http.Serve(l, http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/" {
			//fmt.Println("HandlerFunc1!!")
			w.Write([]byte(`<script>location.href = "/close?" + location.hash.substring(1);</script>`))
		} else {
			//fmt.Println("HandlerFunc2!!")
			w.Write([]byte(`<script>window.open("about:blank","_self").close()</script>`))
			w.(http.Flusher).Flush()
			quit <- req.URL.Query().Get("access_token")
		}
	}))

	token = <-quit
	err = save(token, &now, filename)
	if err != nil {
		fmt.Printf("save index error: %v\n", err)
		os.Exit(1)
	}
	return token, nil
}
