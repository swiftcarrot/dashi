GOFMT_FILES?=$$(find . -name '*.go')

all: fmt build

fmt:
	gofmt -w $(GOFMT_FILES)

build:
	find . -name ".DS_Store" -delete
	go mod tidy
	~/go/bin/packr2
	go build -ldflags "-X github.com/swiftcarrot/dashi/cmd.Version=`git rev-parse HEAD`" -o bin/dashi
	go install

clean:
	rm bin/dashi
	~/go/bin/packr2 clean
