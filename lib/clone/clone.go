package clone

import (
	"os/exec"
	"github.com/libgit2/git2go"
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
