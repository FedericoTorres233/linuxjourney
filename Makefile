# Default target
all: build

build-app:
	@echo "Building application!"
	go build -o bin/linuxjourney main.go

run: build
	./bin/linuxjourney

test:
	@echo "Running tests!"
	go test -v ./... -count=1

tidy:
	go mod tidy -v

format:
	@echo "Beautify code"
	go run mvdan.cc/gofumpt@latest -w -l .

release:
	@echo "Generating a new release!"
	$(MAKE) build-app
	cp -r ./public/* ./bin
	tar cvfz linuxjourney-release.tar.gz ./bin
	rm ./bin/*

build:
	@echo "Building application in dev mode!"
	$(MAKE) build-app

clean:
	@echo "Cleaning up"
	rm -rf ./bin
