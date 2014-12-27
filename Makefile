GOBIN = $(shell which go)
GIT = $(shell which git)

.PHONY: configure, install


configure: 
	$(GOBIN) get
	

install:
	$(GOBIN) build -o dist/mogoint
	$(shell mv dist/mogoint /usr/local/bin/mogoint)

