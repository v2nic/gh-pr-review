#!/bin/bash
set -e

VERSION=${1#v}

mkdir -p dist

# Build for Windows
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o dist/gh-pr-review_${VERSION}_windows-amd64.exe .

# Build for macOS
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o dist/gh-pr-review_${VERSION}_darwin-amd64 .
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o dist/gh-pr-review_${VERSION}_darwin-arm64 .

# Build for Linux
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o dist/gh-pr-review_${VERSION}_linux-amd64 .
GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o dist/gh-pr-review_${VERSION}_linux-arm64 .
