.PHONY = installation

installation:
	go get -u github.com/ipfs/go-ipfs-api

server:
	@ipfs daemon

add:
	@go run addFile/main.go

get:
	@go run getFile/main.go