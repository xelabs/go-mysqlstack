export PATH := $(GOPATH)/bin:$(PATH)

fmt:
	go fmt ./...

test:
	@echo "--> Testing..."
	@$(MAKE) testsqlparser
	@$(MAKE) testcommon
	@$(MAKE) testproto
	@$(MAKE) testpacket
	@$(MAKE) testdriver

testsqlparser:
	go test -v ./sqlparser
testcommon:
	go test -v ./common
testproto:
	go test -v ./proto
testpacket:
	go test -v ./packet
testdriver:
	go test -v ./driver

COVPKGS = ./sqlparser ./common ./proto ./packet ./driver
coverage:
	go get github.com/pierrre/gotestcover
	gotestcover -coverprofile=coverage.out -v $(COVPKGS)
	go tool cover -html=coverage.out

.PHONY: fmt testcommon testproto testpacket testdriver coverage
