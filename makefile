fmt:
	go fmt ./...

test:
	@echo "--> Testing..."
	@$(MAKE) testcommon
	@$(MAKE) testproto
	@$(MAKE) testpacket

testcommon:
	go test -v ./common

testproto:
	go test -v ./proto

testpacket:
	go test -v ./packet

COVPKGS = ./common ./proto ./packet
coverage:
	gotestcover -coverprofile=coverage.txt -v $(COVPKGS)

.PHONY: fmt testcommon testproto testpacket coverage
