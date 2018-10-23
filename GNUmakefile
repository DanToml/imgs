CIRCLE_BUILD_NUM ?= 0
TAG?=0.0.$(CIRCLE_BUILD_NUM)-$(shell git rev-parse --short HEAD)

DATE = $(shell date "+%FT%T%z")

PREFIX?=$(shell pwd)
NAME := imgs
PKG := github.com/dantoml/imgs/api

BUILDTAGS ?=

# Set the build dir, where built cross-compiled binaries will be output
BUILDDIR := ${PREFIX}/cross

# Populate version variables
# Add to compile time flags
CTIMEVAR=-X $(PKG)/cmd.BuildDate=$(DATE) -X $(PKG)/cmd.VERSION=$(TAG)
GO_LDFLAGS=-ldflags "-w $(CTIMEVAR)"

# List the GOOS and GOARCH to build
GOOSARCHES = darwin/amd64 darwin/386 freebsd/amd64 freebsd/386 linux/arm linux/arm64 linux/amd64 linux/386 windows/amd64

.PHONY: tools
tools:
	@echo "+ $@"
	go get -u gotest.tools/gotestsum
	go get -u github.com/alecthomas/gometalinter

.PHONY: build
build: dist/$(NAME) ## Builds a dynamic executable or package

.PHONY: dist/$(NAME)
dist/$(NAME):
	@echo "+ $@"
	go build -tags "$(BUILDTAGS)" ${GO_LDFLAGS} -o $@ .

