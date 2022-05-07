package main

import (
	"bytes"
	"encoding/json"
	"log"
)

func JsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", " ")
	if err != nil {
		return in
	}
	return out.String()
}

func main() {
	c := ClientConnection()
	// result := c.GetRepoInfo("ci-bot")
	result := c.GetPullRequest("cloud-estate", 61)
	output := JsonPrettyPrint(result)
	log.Println(output)
}
