GOFMT_FILES?=$$(find . -name '*.go')

all: fmt build install

fmt:
	gofmt -w $(GOFMT_FILES)
	yarn prettier --write '**/*.{js,json,jsx,ts,tsx,graphql,yaml,yml,md}' --loglevel warn

build:
	go mod tidy
	go build -ldflags "-X github.com/swiftcarrot/dashi/cmd.Version=`git rev-parse HEAD`" -o bin/dashi

install: build
	install bin/dashi /usr/local/bin

clean:
	rm bin/dashi
