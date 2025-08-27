package main

import (
	"flag"
	"fmt"

	"github.com/muromeo1/go/pkg/prdescription"
)

const (
	url = "https://api.openai.com/v1/responses"
)

func main() {
	key := flag.String("open-ai-key", "", "Open AI key")
	branch := flag.String("branch", "main", "Branch to compare against")

	flag.Parse()
	if *key == "" {
		fmt.Printf("Open ai key must be provided")
		return
	}

	git := prdescription.NewGitFetcher(*branch)
	gitLog := git.Log()

	client := prdescription.NewClient(url, *key, "gpt-4.1-nano")
	resp := client.Responses(gitLog)

	prdescription.CopyToClipboard(resp)
	fmt.Println(resp)
}
