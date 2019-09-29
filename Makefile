default: build

build:
	go build -o bin/encrypt confidentiality/cmd/encrypt/encrypt.go
	go build -o bin/decrypt confidentiality/cmd/decrypt/decrypt.go
