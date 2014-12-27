# Mogo Installer

Install MogoCMS and dependencies

This will do the following;

- Install Mongodb, and PHP (if needed) `--deps true`
  - with brew for osx `--deps true --osx`
  - with yum for RHEL/CentOS `--deps true --rhel`
  - with apt-get for debian `--deps true --deb`
- Get the latest stable version of MogoCMS `--HEAD true`
- Check into VCS `--vcs true`

## Usage

1. Download the latest installer [From Here](https://github.com/MogoCMS/mogo-installer/releases)
2. run it from the terminal with the appropriate flags read above

## Building From Source

1.     $ git clone https://github.com/MogoCMS/mogo-installer && cd mogo-Installer
2.     $ make configure
3.     $ make install
4.     $ mogoint --help #done

