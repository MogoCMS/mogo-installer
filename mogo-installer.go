package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Mogo Installer"
	app.Usage = "Install and Setup MogoCMS and its Dependencies"

	// Setup Flags and Arguments
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "deps",
			Value: " ",
			Usage: "Install Dependencies",
		},
		cli.StringFlag{
			Name:  "HEAD",
			Value: " ",
			Usage: "Install from HEAD (Master Branch) (Unstable)",
		},
	}

	app.action = func(c *cli.Context) {
		// Check 'deps' flag
		if c.String("deps") == "true" {
			println("Installing Dependencies")
		}
		if c.String("deps") == " " {
			println("Skipping dependency installation")
		}
		// Check 'HEAD' flag
		if c.String("HEAD") == "true" {
			println("Installing MogoCMS from HEAD")
		}
		if c.String("HEAD") == " " {
			println("Installing MogoCMS latest stable")
		}
	}

	app.Run(os.Args)
}
