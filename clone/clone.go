package clone

import (
	"github.com/codeskyblue/go-sh"
)

func cloneRepo(repo string) {
	sh.Command("git", "clone %v", repo)
	return
}
