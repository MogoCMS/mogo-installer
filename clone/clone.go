package clone

import (
	"fmt"
	"log"
	"os/exec"
)

func clone(repo String) {
	cmd := exec.Command("git", "clone", repo)
	printCommand(cmd)
	output, err := cmd.CombinedOutput()
	printError(err)
	printOutput(output)
}
