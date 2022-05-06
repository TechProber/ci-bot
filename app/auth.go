package main

import (
	"context"
	"net/url"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func ClientConnection() *ClientService {
	token := ""

	ctx := context.Background()
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tokenClient := oauth2.NewClient(ctx, tokenService)

	client := github.NewClient(tokenClient)
	client.BaseURL, _ = url.Parse(ApiURL)

	return &ClientService{
		Context: ctx,
		Client:  client,
	}

}
