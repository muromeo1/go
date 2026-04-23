package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/muromeo1/go/pkg/copilotusage"
)

const (
	defaultToken = "<token>"
	url          = "https://api.github.com/copilot_internal/user"
)

func main() {
	refreshRate := flag.Int("refresh", 10, "Refresh rate in seconds")
	copilotToken := flag.String("token", "", "GitHub Copilot API token")
	flag.Parse()

	var token string
	if *copilotToken == "" {
		token = defaultToken
	}

	tick := time.Tick(time.Second * time.Duration(*refreshRate))
	printUsage(token)

	for range tick {
		printUsage(token)
	}
}

func printUsage(token string) {
	usage := copilotusage.FetchUsage(url, token)

	fmt.Println(usage)
}
