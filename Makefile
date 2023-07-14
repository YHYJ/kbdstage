export GO15VENDOREXPERIMENT=1

SHA=$(shell git rev-parse --short HEAD)
COUNT=$(shell git rev-list --count HEAD)

BUILDTAG=${COUNT}.${SHA}

BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
ifeq ($(BRANCH),main)
BUILDTYPE=release
else
BUILDTYPE=$(BRANCH)
endif

all: deps bundle build

build: bundle
	@go build -trimpath -ldflags \
		"-s -w -X main.Build=${BUILDTAG} -X main.Type=${BUILDTYPE}" \
		-o kbdstage

bundle: deps
	@bash load_ttf.sh

clean:
	@rm -f kbdstage function/resource_ttf.go

tidy:
	@echo "Tidying up dependencies..."
	@go mod tidy

deps:
	@echo "Getting required dependencies..."
	@go install github.com/kevinburke/go-bindata/...

.PHONY: build deps bundle
