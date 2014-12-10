package main

import (
	"github.com/codegangsta/cli"
	"github.com/mogocms/mogo-installer/clone"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Mogo Installer"
	app.Usage = "Install and Setup MogoCMS and its Dependencies"

	// Git Stuff
	headRepo := ("github.com/mogocms/mogo/tree/master")
	stableRepo := ("github.com/mogocms/mogo/tree/stable")
	cl := clone.Clone

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
		cli.StringFlag{
			Name:  "conf",
			Value: " ",
			Usage: "write the config file",
		},
		cli.StringFlag{
			Name:  "vcs",
			Value: " ",
			Usage: "Check MogoCMS into git",
		},
	}

	app.Action = func(c *cli.Context) {
		// Check 'deps' flag
		if c.String("deps") == "true" {
			println("Installing dependencies")
		}
		if c.String("deps") == " " {
			println("Skipping dependency installation")
		}
		// Check 'HEAD' flag
		if c.String("HEAD") == "true" {
			println("Installing MogoCMS from HEAD")
			cl(headRepo)
		}
		if c.String("HEAD") == " " {
			println("Installing MogoCMS latest stable")
			cl(stableRepo)
		}
		// Check 'conf' flag
		if c.String("conf") == "true" {
			println("Writing config file")
		}
		if c.String("conf") == " " {
			println("Skipping configuration file, please write it manually")
		}
		// Check 'vcs' flag
		if c.String("vcs") == "true" {
			println("Checking your new MogoCMS installation into git")
		}
		if c.String("vcs") == " " {
			println("Skipping VCS")
		}
	}

	app.Run(os.Args)
}
