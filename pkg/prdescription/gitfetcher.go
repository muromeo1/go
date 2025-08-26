package prdescription

import (
	"os/exec"
)

type GitFetcher struct {
	Branch string
}

func NewGitFetcher(branch string) *GitFetcher {
	return &GitFetcher{
		Branch: branch,
	}
}

func (g *GitFetcher) Diff() string {
	cmd := exec.Command("git", "diff", g.Branch)
	output, _ := cmd.Output()

	return string(output)
}
