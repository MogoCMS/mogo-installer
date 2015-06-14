package main

import (
	"github.com/codegangsta/cli"
	"github.com/MogoCMS/mogo-installer/lib/clone"
	"github.com/MogoCMS/mogo-installer/lib/cmd"
	"os"
	"fmt"
	"github.com/docopt/docopt-go"
	"runtime"
)

// func main() {
// 	app := cli.NewApp()
// 	app.Name = "Mogo Installer"
// 	app.Usage = "Install and Setup MogoCMS and its Dependencies"

// 	// Git Clone
// 	cl := clone.Clone
// 	runcmd := cmd.Cmd

// 	// Setup Flags and Arguments
// 	app.Flags = []cli.Flag{
// 		cli.StringFlag{
// 			Name:  "deps",
// 			Value: " ",
// 			Usage: "Install Dependencies",
// 		},
// 		cli.StringFlag{
// 			Name:  "os",
// 			Value: " ",
// 			Usage: "Specify what system you are installing mogo on",
// 		},
// 		cli.StringFlag{
// 			Name:  "HEAD",
// 			Value: " ",
// 			Usage: "Install from HEAD (Master Branch) (Unstable)",
// 		},
// 		cli.StringFlag{
// 			Name:  "vcs",
// 			Value: " ",
// 			Usage: "Check MogoCMS into git",
// 		},
// 	}

// 	app.Action = func(c *cli.Context) {
// 		// Check 'deps' flag
// 		if c.String("deps") == "true" {
// 			println("Installing dependencies")
// 			if c.String("os") == "osx" {
// 				println("Installing on OSX (using brew)")
// 				runcmd("brew install mongodb")
// 				runcmd("brew tap Homebrew/homebrew-php")
// 				runcmd("brew install php55")
// 				println("installed all dependencies for MogoCMS")
// 			}
// 			if c.String("os") == "deb" {
// 				println("Installing on debian (using apt-get)")
// 				runcmd("sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 7F0CEB10")
// 				runcmd("echo 'deb http://downloads-distro.mongodb.org/repo/ubuntu-upstart dist 10gen' | sudo tee /etc/apt/sources.list.d/mongodb.list")
// 				runcmd("sudo apt-get update")
// 				runcmd("sudo apt-get install -y mongodb-org")
// 				runcmd("sudo apt-get install php5 php-pear")
// 				println("Installed all depenencies for mogoCMS")
// 			}
// 			if c.String("os") == "rhel" {
// 				println("Installing on RHEL/CentOS (using yum)")
// 				runcmd("echo \"[mongodb]\" >> /etc/yum.repos.d/mongodb.repo")
// 				runcmd("echo \"name=MongoDB Repository\" >> /etc/yum.repos.d/mongodb.repo")
// 				runcmd("echo \"baseurl=http://downloads-distro.mongodb.org/repo/redhat/os/x86_64/\" >> /etc/yum.repos.d/mongodb.repo")
// 				runcmd("echo \"gpgcheck=0\" >> /etc/yum.repos.d/mongodb.repo")
// 				runcmd("echo \"enabled=1\" >> /etc/yum.repos.d/mongodb.repo")
// 				runcmd("sudo yum install -y mongodb-org")
// 				runcmd("yum install php55w php55w-opcache")
// 			}
// 		}
// 		if c.String("deps") == " " {
// 			println("Skipping dependency installation")
// 		}
// 		// Check 'HEAD' flag
// 		if c.String("HEAD") == "true" {
// 			println("Installing MogoCMS from HEAD")
// 			cl("master", "https://github.com/mogocms/mogo.git")
// 		}
// 		if c.String("HEAD") == " " {
// 			println("Installing MogoCMS latest stable")
// 			cl("stable", "https://github.com/mogocms/mogo.git")
// 		}
// 		runcmd("cp ./mogo/install.sample.php ./mogo/install.php")
// 		// Check 'vcs' flag
// 		if c.String("vcs") == "true" {
// 			println("Checking your new MogoCMS installation into git")
// 			runcmd("git init")
// 			runcmd("git add .")
// 			runcmd("git commit -am 'Initial installation of MogoCMS'")
// 		}
// 		if c.String("vcs") == " " {
// 			println("Skipping VCS")
// 		}
// 	}

// 	app.Run(os.Args)
// }

func main() {
	usage := `Usage:
mogo-installer [--deps=<os>]

Options:
  -h, --help              Show this screen.
  -v, --version           Show version.
  -d, --dependencies      Install dependencies for your operating system
  `
	args, _ := docopt.Parse(usage, nil, true, "Mogo Installer", false)
    
    Dep, DepSet := args['--dependencies'].(string)

    if DepSet {
 		if Dep = "deb" {
 			println("Installing on debian (using apt-get)")
			runcmd("sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 7F0CEB10")
			runcmd("echo 'deb http://downloads-distro.mongodb.org/repo/ubuntu-upstart dist 10gen' | sudo tee /etc/apt/sources.list.d/mongodb.list")
			runcmd("sudo apt-get update")
			runcmd("sudo apt-get install -y mongodb-org")
			runcmd("sudo apt-get install php5 php-pear")
			println("Installed all depenencies for mogoCMS")
 		}
 		if Dep = "osx" {
 			println("Installing on OSX (using brew)")
			runcmd("brew install mongodb")
			runcmd("brew tap Homebrew/homebrew-php")
			runcmd("brew install php55")
			println("installed all dependencies for MogoCMS")
 		}
 		if Dep = "rhel" {
			println("Installing on RHEL/CentOS (using yum)")
			runcmd("echo \"[mongodb]\" >> /etc/yum.repos.d/mongodb.repo")
			runcmd("echo \"name=MongoDB Repository\" >> /etc/yum.repos.d/mongodb.repo")
			runcmd("echo \"baseurl=http://downloads-distro.mongodb.org/repo/redhat/os/x86_64/\" >> /etc/yum.repos.d/mongodb.repo")
			runcmd("echo \"gpgcheck=0\" >> /etc/yum.repos.d/mongodb.repo")
			runcmd("echo \"enabled=1\" >> /etc/yum.repos.d/mongodb.repo")
			runcmd("sudo yum install -y mongodb-org")
			runcmd("yum install php55w php55w-opcache")
 		}
    }
}