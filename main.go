package main

import (
	"fmt"
	"time"
	"net/http"
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
	httpVerb func(string) (*http.Response, error),
	url string,

) {
	resp, err := cb(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func main() {
	netClient := newClient()
	s := new(Wallpapers)

	fmt.Println(netClient, s)
}
