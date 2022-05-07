package main

import (
	"encoding/json"
	"log"
)

func (c *ClientService) GetRepoInfo(repository string) string {
	result, _, err := c.Client.Repositories.Get(c.Context, RepoOwner, repository)
	if err != nil {
		log.Fatalf("Error occurred while fetching API data %v\n", err)
	}

	bytes, _ := json.Marshal(&Repository{
		FullName:    *result.FullName,
		Description: *result.Description,
		HTMLURL:     *result.HTMLURL,
		URL:         *result.URL,
		UpdatedAt:   *result.UpdatedAt,
	})

	return string(bytes)
}

func (c *ClientService) GetPullRequest(repository string, number int) string {
	result, _, err := c.Client.PullRequests.Get(c.Context, RepoOwner, repository, number)
	if err != nil {
		log.Fatalf("Error occurred while fetching API data %v\n", err)
	}

	bytes, _ := json.Marshal(&PullRequest{
		ID:        *result.ID,
		Number:    *result.Number,
		State:     *result.State,
		Title:     *result.Title,
		URL:       *result.URL,
		Assignees: *&result.Assignees,
	})

	return string(bytes)
}
