package main

import (
	"encoding/json"
	"log"
)

func (c *ClientService) GetRepoInfo(repository string) string {
	result, _, err := c.Client.Repositories.Get(c.Context, "TechProber", repository)
	if err != nil {
		log.Fatalf("Error occurred while fetching API data %v\n", err)
	}

	bytes, _ := json.Marshal(Repository{
		FullName:    *result.FullName,
		Description: *result.Description,
		HTMLURL:     *result.HTMLURL,
		URL:         *result.URL,
		UpdatedAt:   *result.UpdatedAt,
	})

	return string(bytes)
}
