package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
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

func main() {
	netClient := newClient()
	walls := new(Wallpapers)

	resp, err := netClient.Get("https://wallhaven.cc/api/v1/search")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(walls); err != nil {
		panic(err)
	}

	for _, wall := range walls.Data {
		startAt := strings.LastIndex(wall.Path, wall.ID)
		filename := wall.Path[startAt:]

		file, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		resp, err := netClient.Get(wall.Path)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		if _, err := io.Copy(file, resp.Body); err != nil {
			panic(err)
		}
	}
}
