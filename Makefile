GOCMD=go

.PHONY: all test

all:
	$(GOCMD) version

test:
	$(GOCMD) test -v ./...