package main

import "os"

const ApiURL = "https://api.github.com/"
const AppID = 198166
const RepoOwner = "TechProber"
const InstallationID = 25506861

func GetKey() string {
	if os.Getenv("DEBUG") == "TRUE" {
		return "/home/kev/.local/github/ci-bot/techprober-ci-bot.private-key.pem"
	} else {
		return "/private-key.pem"
	}
}
