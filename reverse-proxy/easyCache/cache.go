package easyCache

import (
	"encoding/gob"
	"os"
)

type httpCache struct {
	filepath string
	items    map[string][]byte
}

// Current HttpCache implementation is NOT "thread-safe".
// use goroutine and channel for singleton
// getReq, getRes
// setReq, delReq, loadReq, saveReq, resChan
// itemsReq, itemsRes

func NewHttpCache(cacheFilename string) httpCache {
	c := httpCache{
		filepath: cacheFilename,
		items:    make(map[string][]byte),
	}
	return c
}

func (c *httpCache) Items() map[string][]byte {
	return c.items
}

func (c *httpCache) Set(key string, value []byte) {
	c.items[key] = value
}

func (c *httpCache) Del(key string) {
	delete(c.items, key)
}

func (c *httpCache) Get(key string) ([]byte, bool) {
	val, ok := c.items[key]
	return val, ok
}

func (c *httpCache) Load() error {
	fd, err := os.Open(c.filepath)
	if err != nil {
		return err
	}
	defer fd.Close()

	dec := gob.NewDecoder(fd)
	err = dec.Decode(&c.items)
	if err != nil {
		return err
	}
	return nil
}

func (c *httpCache) Save() error {
	fd, err := os.Create(c.filepath)
	if err != nil {
		return err
	}
	defer fd.Close()

	enc := gob.NewEncoder(fd)
	err = enc.Encode(c.items)
	if err != nil {
		return err
	}
	return nil
}
