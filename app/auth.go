package main

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/github"
)

func ClientConnection() *ClientService {
	transport := http.DefaultTransport
	install, err := ghinstallation.NewKeyFromFile(transport, AppID, InstallationID, GetKey())
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	client := github.NewClient(&http.Client{
		Transport: install,
		Timeout:   5 * time.Second,
	})

	client.BaseURL, _ = url.Parse(ApiURL)

	return &ClientService{
		Context: ctx,
		Client:  client,
	}
}
