export PATH := $(GOPATH)/bin:$(PATH)

fmt:
	go fmt ./...

test:
	@echo "--> Testing..."
	@$(MAKE) testcommon
	@$(MAKE) testproto
	@$(MAKE) testpacket
	@$(MAKE) testdriver

testcommon:
	go test -v ./common

testproto:
	go test -v ./proto

testpacket:
	go test -v ./packet

testdriver:
	go test -v ./driver

COVPKGS = ./common ./proto ./packet
coverage:
	gotestcover -coverprofile=coverage.txt -v $(COVPKGS)

.PHONY: fmt testcommon testproto testpacket coverage
