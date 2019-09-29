build:
	go build -o bin/encrypt cmd/encrypt/encrypt.go
	go build -o bin/decrypt cmd/decrypt/decrypt.go
	go build -o bin/pkencrypt cmd/pkencrypt/pkencrypt.go
	go build -o bin/pkdecrypt cmd/pkdecrypt/pkdecrypt.go
	go build -o bin/hash cmd/hash/hash.go
	go build -o bin/pksign cmd/pksign/pksign.go
	go build -o bin/pkverify cmd/pkverify/pkverify.go
