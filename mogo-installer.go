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
	}

	app.action = func(c *cli.Context) {
		// Runtime
		if c.String("deps") == "true" {
			println("Installing Dependencies")
		}
		if c.String("deps") == " " {
			println("Skipping Dependency Installation")
		}
	}

	app.Run(os.Args)
}
