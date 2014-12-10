package clone

import (
	"github.com/codeskyblue/go-sh"
)

func Clone(repo string) {
	sh.Command("git", "clone %v", repo)
	return
}
