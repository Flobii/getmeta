.PHONY: example
.PHONY: install

example:
	@echo "\044 go run getmeta.go http://neuralnetworksanddeeplearning.com/"
	@go run getmeta.go http://neuralnetworksanddeeplearning.com/

build:
	@echo Building...
	@go build getmeta.go
	@echo Done!

compile:
	@echo Compiling...
	@echo Windows...
	@GOOS=windows GOARCH=amd64 go build -o bin/getmeta-amd64.exe getmeta.go
	@GOOS=windows GOARCH=386 go build -o bin/getmeta-386.exe getmeta.go
	@echo Linux...
	@GOOS=linux GOARCH=amd64 go build -o bin/getmeta-amd64-linux getmeta.go
	@GOOS=linux GOARCH=386 go build -o bin/getmeta-386-linux getmeta.go
	@echo MacOS...
	@GOOS=darwin GOARCH=amd64 go build -o bin/getmeta-amd64-darwin getmeta.go
	@echo Done!

install:
	@echo Installing...
	@go insall getmeta.go
	@echo Done!
