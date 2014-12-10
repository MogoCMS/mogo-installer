package cmd

import (
	"fmt"
	"os/exec"
	"strings"
)

func Cmd(cmd string) {
	fmt.Println(cmd)
	parts := strings.Fields(cmd)
	out, err := exec.Command(parts[0], parts[1]).Output()
	if err != nil {
		fmt.Println("error occured")
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)
	return
}
