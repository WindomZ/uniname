BINARY := uniname

VERSION := $(shell git describe --tags --abbrev=0 2>/dev/null || (git describe --always --long --dirty | tr '\n' '-';date +%Y.%m.%d))
LDFLAGS = -ldflags "-w -s -X main.version=${VERSION}"
LDFLAGS_DEV = -ldflags "-X main.version=${VERSION}"

GOX := $(shell command -v gox 2> /dev/null)

.PHONY: all

gox:
ifndef GOX
	$(error "gox is not available. Please install from https://github.com/mitchellh/gox")
endif

#Build release builds
release: gox
	@gox -osarch="darwin/386 darwin/amd64 linux/386 linux/amd64 windows/386 windows/amd64" ${LDFLAGS} -output="bin/{{.Dir}}_{{.OS}}_{{.Arch}}"

#Build a development build
build:
	@go build ${LDFLAGS_DEV} -o bin/${BINARY}

#Install a release build on your local system
install: clean
	@go install ${LDFLAGS}

clean:
	@rm -rf bin/*
	@go clean -i
