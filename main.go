package main

import (
	"github.com/codegangsta/cli"
	"github.com/MogoCMS/mogo-installer/lib/clone"
	"github.com/MogoCMS/mogo-installer/lib/cmd"
	"os"
	"fmt"
	"github.com/docopt/docopt-go"
)

// Git Clone
cl := clone.Clone
runcmd := cmd.Cmd

func main() {
	usage := `Usage:
mogo-installer [--deps=<os>] [--HEAD] [--vcs]

Options:
  -h, --help              Show this screen.
  -v, --version           Show version.
  -d, --dependencies      Install dependencies for your operating system
  -h, --HEAD              Optionally get the HEAD version
  `
	args, _ := docopt.Parse(usage, nil, true, "Mogo Installer", false)
    
    Dep, DepSet := args['--dependencies'].(string)
    headFlag := args['--HEAD']

    // Check Dependency flag
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
 		else if Dep = "osx" {
 			println("Installing on OSX (using brew)")
			runcmd("brew install mongodb")
			runcmd("brew tap Homebrew/homebrew-php")
			runcmd("brew install php55")
			println("installed all dependencies for MogoCMS")
 		}
 		else if Dep = "rhel" {
			println("Installing on RHEL/CentOS (using yum)")
			runcmd("echo \"[mongodb]\" >> /etc/yum.repos.d/mongodb.repo")
			runcmd("echo \"name=MongoDB Repository\" >> /etc/yum.repos.d/mongodb.repo")
			runcmd("echo \"baseurl=http://downloads-distro.mongodb.org/repo/redhat/os/x86_64/\" >> /etc/yum.repos.d/mongodb.repo")
			runcmd("echo \"gpgcheck=0\" >> /etc/yum.repos.d/mongodb.repo")
			runcmd("echo \"enabled=1\" >> /etc/yum.repos.d/mongodb.repo")
			runcmd("sudo yum install -y mongodb-org")
			runcmd("yum install php55w php55w-opcache")
 		} else {
 			fmt.Println("Please provide an OS to install for either deb, osx, or rhel")
 		}
    }

    // Clone the MogoCMS From master/stable
    if headFlag {
    	fmt.Println('Cloning from master branch (WARNING: May be unstable)')
    	cl("master", "https://github.com/mogocms/mogo.git")
    } else {
    	fmt.Println('Cloning from stable branch')
    	cl("stable", "https://github.com/mogocms/mogo.git")
    }

    // Create the config install file
    runcmd("cp ./mogo/install.sample.php ./mogo/install.php")
}