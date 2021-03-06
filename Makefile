GOBIN = $(shell which go)
GIT = $(shell which git)

.PHONY: configure, install, build

configure: 
	$(GOBIN) get github.com/codegangsta/cli

build:
	$(GOBIN) build -o dist/mogoint

install:
	$(shell make build)
	$(shell mv dist/mogoint /usr/local/bin/mogoint)
