# ******************************************************
# Author       	:	serialt 
# Email        	:	tserialt@gmail.com
# Filename     	:   Makefile
# Version      	:	v1.3.0
# Created Time 	:	2021-06-25 10:47
# Last modified	:	2023-03-12 22:47
# By Modified  	: 
# Description  	:       build go package
#  
# ******************************************************


PROJECT_NAME= cli


GOBASE=$(shell pwd)
GOFILES=$(wildcard *.go)


BRANCH := $(shell git symbolic-ref HEAD 2>/dev/null | cut -d"/" -f 3)
# BRANCH := `git fetch --tags && git tag | sort -V | tail -1`
# BUILD := $(shell git rev-parse --short HEAD)
BUILD_DIR := $(GOBASE)/dist
VERSION = $(BRANCH)

BuildTime := $(shell date -u  '+%Y-%m-%d %H:%M:%S %Z')
GitHash := $(shell git rev-parse HEAD)
GoVersion := $(shell go version | awk '{print $3}')
Maintainer := tserialt@gmail.com 

PKGFLAGS := " -s -w -X 'main.APPVersion=$(VERSION)'  -X 'main.BuildTime=$(BuildTime)' -X 'main.GitCommit=$(GitHash)' "

APP_NAME = $(PROJECT_NAME)
# go-pkg.v0.1.1-linux-amd64

.PHONY: clean
clean:
	\rm -rf dist/$(PROJECT_NAME)* 

.PHONY: serve
serve:
	go run .

.PHONY: build
build: clean
	go build -ldflags $(PKGFLAGS) -o "dist/$(APP_NAME)" cmd/sugar.go
	@echo "编译完成"

.PHONY: release
release: clean
	go mod tidy
	GOOS="windows" GOARCH="amd64" go build -ldflags $(PKGFLAGS) -v -o "sugar/$(APP_NAME)-windows-amd64.exe"  cmd/sugar.go
	GOOS="linux"   GOARCH="amd64" go build -ldflags $(PKGFLAGS) -v -o "sugar/$(APP_NAME)-linux-amd64"        cmd/sugar.go
	GOOS="linux"   GOARCH="arm64" go build -ldflags $(PKGFLAGS) -v -o "sugar/$(APP_NAME)-linux-arm64"        cmd/sugar.go
	GOOS="darwin"  GOARCH="amd64" go build -ldflags $(PKGFLAGS) -v -o "sugar/$(APP_NAME)-darwin-amd64"       cmd/sugar.go
	GOOS="darwin"  GOARCH="arm64" go build -ldflags $(PKGFLAGS) -v -o "sugar/$(APP_NAME)-darwin-arm64"       cmd/sugar.go
	@echo "******************"
	@echo " release succeed "
	@echo "******************"
	ls -la dist/$(PROJECT_NAME)*