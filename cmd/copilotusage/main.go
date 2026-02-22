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
	refreshRate := flag.Int("refresh", 5, "Refresh rate in seconds")
	token := flag.String("token", defaultToken, "GitHub Copilot API token")
	flag.Parse()

	tick := time.Tick(time.Second * time.Duration(*refreshRate))
	printUsage(*token)

	for range tick {
		printUsage(*token)
	}
}

func printUsage(token string) {
	usage := copilotusage.FetchUsage(url, token)

	fmt.Println(usage)
}
