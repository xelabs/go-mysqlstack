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

.PHONY: fmt testcommon testproto testpacket
