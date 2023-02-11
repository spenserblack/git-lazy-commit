#!/bin/sh
# Builds the binary for several platforms
set -e

build() {
	platform=$1
	goos=$(echo $platform | cut -d- -f1)
	goarch=$(echo $platform | cut -d- -f2)
	ext=""
	if [ "$goos" = "windows" ]; then
		ext=".exe"
	fi
	GOOS=$goos GOARCH=$goarch go build -o "dist/git-lzc-$platform$ext" ./cmd/git-lzc
}

build "darwin-amd64"
build "linux-amd64"
build "linux-386"
build "windows-amd64"
build "windows-386"
build "openbsd-amd64"
build "openbsd-386"
build "freebsd-amd64"
build "freebsd-386"
build "netbsd-amd64"
build "netbsd-386"
