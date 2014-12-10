package clone

import (
	"os/exec"
)

func Clone(Branch string, Repo string) {

	cmd := exec.Command("git", "clone", "-b", Branch, Repo)
	out, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	println(string(out))
}
