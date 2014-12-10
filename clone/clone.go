package clone

import (
	"fmt"
	"github.com/codeskyblue/go-sh"
)

func Clone(repo string, branch string) {
	clonecmd := fmt.Sprintf("git clone -b %s %s", branch, repo)
	sh.Command(clonecmd).Run()
}
