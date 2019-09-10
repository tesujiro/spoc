package api

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/tesujiro/spoc/global"
)

const spotify_base_url = "https://api.spotify.com"

type Api struct {
	token    string
	Base_url string //TODO:
}

func New() *Api {
	token, err := getAccessToken() // CAUTION: The access tokens expire after 1 hour.
	if err != nil {
		log.Fatal("faild to get access token:", err)
	}
	return &Api{
		token:    token,
		Base_url: spotify_base_url,
	}
}

/*
type SpotifyAPI struct {
	cmd      string
	target   string
	usage    string
	desc     string
	endpoint string
}
*/

func (api *Api) Get(endpoint string, params url.Values) ([]byte, error) {
	return api.call("GET", endpoint, params, nil)
}

func (api *Api) Put(endpoint string, params url.Values, body io.Reader) ([]byte, error) {
	return api.call("PUT", endpoint, params, body)
}

func (api *Api) Post(endpoint string, params url.Values, body io.Reader) ([]byte, error) {
	return api.call("POST", endpoint, params, body)
}

func (api *Api) call(method, endpoint string, params url.Values, body io.Reader) ([]byte, error) {
	if os.Getenv("ReverseProxy") != "" {
		proxy := os.Getenv("ReverseProxy")
		endpoint = strings.Replace(endpoint, api.Base_url, proxy, 1)
	}

	//fmt.Println("endpoint:", endpoint)
	baseUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	baseUrl.RawQuery = params.Encode() // Escape Query Parameters
	if global.FlagRawJson {
		fmt.Println(baseUrl)
	}

	req, err := http.NewRequest(method, baseUrl.String(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+api.token)

	var client *http.Client
	client = http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("resp=%v\n", resp)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		b, _ := ioutil.ReadAll(resp.Body)
		//fmt.Printf("response.Body=%#v\n", string(b))
		return b, fmt.Errorf("bad response status code %d", resp.StatusCode)
	}
	//fmt.Println("response status code ", resp.StatusCode)

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if global.FlagRawJson {
		fmt.Println(string(b))
	}
	return b, nil
}
