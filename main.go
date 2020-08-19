package main

import (
	"fmt"
	"time"
	"net/http"
)

func newClient() *http.Client {
	netClient := new(http.Client)
	netClient.Timeout = time.Second * 10

	return netClient
}

func main() {
	netClient := newClient()
}
