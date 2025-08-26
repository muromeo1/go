package main

import (
	"flag"
	"fmt"
	"os/exec"

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
	diff := git.Diff()

	client := prdescription.NewClient(url, *key, "gpt-4.1-nano")
	resp := client.Responses(diff)

	copyToClipboard(resp)
	fmt.Println(resp)
}

func copyToClipboard(text string) {
	cmd := exec.Command("pbcopy")
	in, _ := cmd.StdinPipe()
	defer in.Close()

	cmd.Start()
	in.Write([]byte(text))
	in.Close()
	cmd.Wait()

	fmt.Println("Copied to clipboard")
}
