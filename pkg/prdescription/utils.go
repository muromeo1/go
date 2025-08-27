package prdescription

import (
	"fmt"
	"os/exec"
)

func CopyToClipboard(text string) {
	cmd := exec.Command("pbcopy")
	in, _ := cmd.StdinPipe()
	defer in.Close()

	cmd.Start()
	in.Write([]byte(text))
	in.Close()
	cmd.Wait()

	fmt.Println("Copied to clipboard")
}
