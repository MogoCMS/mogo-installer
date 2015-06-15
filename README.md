# Mogo Installer

CLI installer for MogoCMS

## Usage

1. Download the latest installer [From Here](https://github.com/MogoCMS/mogo-installer/releases)
2. run it from the terminal

```bash

mogoint [--deps=<os>] [--HEAD]

Options:
  -h, --help              Show this screen.
  -v, --version           Show version.
  -d, --dependencies      Install dependencies for your operating system
  -H, --HEAD              Optionally get the HEAD version

```

The program is run as `mogoint` and has a few options to customize your installation

* `-d <MY_OS>` - optionally install dependencies PHP5, MongoDB, etc. (its usually a good idea but some people need/want to do it manually) the option takes a paramater which must be either _osx_, _rhel_, or _deb_ the difference is deb uses `apt-get` and rhel uses `yum`.

* `-H` - enabling the `--HEAD` or `-H` flag will tell mogoint to pull in the `master` branch (unstable) instead of the `stable` branch

## Building From Source

1. `$ git clone https://github.com/MogoCMS/mogo-installer && cd mogo-Installer`
2. `$ make configure`
3. `$ make install`
4. `$ mogoint --help #done`

