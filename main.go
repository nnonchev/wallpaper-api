package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Wallpaper struct {
	ID         string `json:"id"`
	Path       string `json:"path"`
	Resolution string `json:"resolution"`
}

type Wallpapers struct {
	Data []Wallpaper `json:"data"`
}

func newClient() *http.Client {
	netClient := new(http.Client)
	netClient.Timeout = time.Second * 10

	return netClient
}

func request(
	httpCb func(string) (*http.Response, error),
	url string,
	cb func(io.Reader, chan Wallpapers),
	chWalls chan Wallpapers,
) {
	resp, err := httpCb(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	cb(resp.Body, chWalls)
}

func main() {
	netClient := newClient()
	s := new(Wallpapers)

	fmt.Println(netClient, s)
}
