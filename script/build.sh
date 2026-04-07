#!/bin/bash
set -e

# Build for Windows (produces .exe)
GOOS=windows GOARCH=amd64 go build -o dist/gh-pr-review_windows_amd64.exe .

# Build for macOS
GOOS=darwin GOARCH=amd64 go build -o dist/gh-pr-review_darwin_amd64 .
GOOS=darwin GOARCH=arm64 go build -o dist/gh-pr-review_darwin_arm64 .

# Build for Linux
GOOS=linux GOARCH=amd64 go build -o dist/gh-pr-review_linux_amd64 .
GOOS=linux GOARCH=arm64 go build -o dist/gh-pr-review_linux_arm64 .
