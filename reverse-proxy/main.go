package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/tesujiro/exp_spotify/reverse-proxy/easyCache"
)

// 参考情報 https://github.com/gregjones/httpcache

const (
	self   = ":8080"
	target = "https://api.spotify.com"
)

var (
	cache         = easyCache.NewHttpCache(cacheFilename)
	cacheFilename = "cache.gob"
)

func main() {
	if len(os.Args) == 2 {
		cacheFilename = os.Args[1]
		fmt.Println("use CacheFile ", os.Args[1])
		cache = easyCache.NewHttpCache(cacheFilename)
	}

	remote, err := url.Parse(target)
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Transport = &myTransport{}
	http.HandleFunc("/", handler(proxy))
	http.HandleFunc("/api", apiHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/save", saveHandler)
	http.HandleFunc("/load", loadHandler)
	err = http.ListenAndServe(self, nil)
	if err != nil {
		panic(err)
	}
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello apiHandler")
	method := r.Method
	v := r.URL.Query()
	fmt.Fprintf(w, "Hello apiHandler: method:%v\n", method)
	var urls []string
	for key, values := range v {
		//fmt.Fprintf(w, "%s = %s\n", key, values)
		if key == "url" {
			urls = values
		}
	}
	/*
		for _, url := range urls {
			fmt.Fprintf(w, "GET key:%v\n", url)
		}
	*/
	switch method {
	case "GET":
	case "POST":
	case "PUT":
	case "DELETE":
		for _, url := range urls {
			if _, ok := cache.Get(url); ok {
				// TODO: lock object
				cache.Del(url)
				fmt.Printf("cache [%v] deleted.\n", url)
				fmt.Fprintf(w, "cache [%v] deleted.\n", url)
			} else {
				fmt.Printf("cache [%v] does not exist.\n", url)
				fmt.Fprintf(w, "cache [%v] does not exist.\n", url)
			}
		}
	default:
		fmt.Fprintf(w, "invalid method:%v\n", method)
	}
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello saveHandler\n")
	// TODO: lock object
	err := cache.Save()
	if err != nil {
		log.Printf("save cache error: %v\n", err)
		return
	}
	for key, _ := range cache.Items() {
		fmt.Fprintf(w, "%v\n", key)
	}
	log.Printf("cache save finished :%v\n", len(cache.Items()))
}

func loadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello loadHandler\n")
	err := cache.Load()
	if err != nil {
		log.Printf("load cache error: %v\n", err)
		return
	}
	for key, _ := range cache.Items() {
		fmt.Fprintf(w, "%v\n", key)
	}
	log.Printf("cache load finished :%v\n", len(cache.Items()))
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello listHandler.\nkeys:\n")
	if len(cache.Items()) == 0 { //TODO
		fmt.Fprintf(w, "no cache key\n")
	} else {
		for key, _ := range cache.Items() {
			fmt.Fprintf(w, "[%v]\n", key)
		}
	}
}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		p.ServeHTTP(w, r)
	}
}

type myTransport struct {
}

func (t *myTransport) CancelRequest(req *http.Request) {
	type canceler interface {
		CancelRequest(*http.Request)
	}
	if cr, ok := http.DefaultTransport.(canceler); ok {
		cr.CancelRequest(req)
	}
}

func (t *myTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	//log.Print("===================REQUEST==========================")
	//log.Println(request.URL)
	//log.Printf("%#v\n", request)
	//log.Print("===================REQUEST==========================")

	cacheKey := request.URL.String()
	var canCache bool = false
	if request.Method == "GET" {
		canCache = true
	}
	if canCache {
		// get cache
		// TODO: lock object
		if cachedBody, ok := cache.Get(cacheKey); ok {
			log.Println("cache hit [" + request.Method + "] " + cacheKey)

			b := bytes.NewBuffer(cachedBody)
			response, err := http.ReadResponse(bufio.NewReader(b), request)
			if err != nil {
				return nil, err
			}
			return response, nil
		}
	}

	// NO CACHE
	log.Println("cache no  [" + request.Method + "] " + cacheKey)
	response, err := http.DefaultTransport.RoundTrip(request)
	//log.Print("===================RESPONSE==========================")
	//log.Printf("%s\n", string(body))
	//log.Print(string(body))
	//log.Print("===================RESPONSE==========================")

	// copying the response body
	body, err := httputil.DumpResponse(response, true)
	if err != nil {
		return nil, err
	}

	// set cache
	if canCache {
		// TODO: lock object
		cache.Set(cacheKey, body)
	}

	return response, err
}
