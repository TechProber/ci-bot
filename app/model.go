package main

import (
	"context"

	"github.com/google/go-github/github"
)

type ClientService struct {
	Context context.Context
	Client  *github.Client
}

type Repository struct {
	FullName    string           `json:"full_name"`
	Description string           `json:"description"`
	HTMLURL     string           `json:"html_url"`
	URL         string           `json:"url"`
	UpdatedAt   github.Timestamp `json:"updatedAt"`
}
