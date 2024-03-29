#!/bin/sh

set -ex

export CGO_ENABLED=0
LDFLAGS="-s -w"

export GOARCH=amd64
go build -trimpath -ldflags "${LDFLAGS}" -o ./release/movie_data_rename-linux-${GOARCH} .

export GOARCH=arm64
go build -trimpath -ldflags "${LDFLAGS}" -o ./release/movie_data_rename-linux-${GOARCH} .
