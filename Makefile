.PHONY: example

example:
	go run getmeta.go http://neuralnetworksanddeeplearning.com/

build:
	@echo Building...
	go build getmeta.go

compile:
	@echo Compiling...
	GOOS=windows GOARCH=amd64 go build -o bin/getmeta-amd64.exe getmeta.go
	GOOS=windows GOARCH=386 go build -o bin/getmeta-386.exe getmeta.go
	GOOS=darwin GOARCH=amd64 go build -o bin/getmeta-amd64-darwin getmeta.go
	GOOS=linux GOARCH=amd64 go build -o bin/getmeta-amd64-linux getmeta.go
	GOOS=linux GOARCH=386 go build -o bin/getmeta-386-linux getmeta.go

install:
	@echo Installing...
	go insall getmeta.go
