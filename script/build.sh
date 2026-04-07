#!/bin/bash
set -e

go build -o gh-pr-review .
zip gh-pr-review-windows-amd64.zip gh-pr-review.exe
GOOS=darwin GOARCH=amd64 go build -o gh-pr-review-darwin-amd64 .
GOOS=darwin GOARCH=arm64 go build -o gh-pr-review-darwin-arm64 .
GOOS=linux GOARCH=amd64 go build -o gh-pr-review-linux-amd64 .
GOOS=linux GOARCH=arm64 go build -o gh-pr-review-linux-arm64 .
zip gh-pr-review-darwin-amd64.zip gh-pr-review-darwin-amd64
zip gh-pr-review-darwin-arm64.zip gh-pr-review-darwin-arm64
zip gh-pr-review-linux-amd64.zip gh-pr-review-linux-amd64
zip gh-pr-review-linux-arm64.zip gh-pr-review-linux-arm64
