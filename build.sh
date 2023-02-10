#!/bin/bash
# Builds the binary for several platforms
# TODO: More portable script that uses sh instead of bash
set -e

platforms=(
	darwin-amd64
	linux-amd64
	linux-386
	windows-amd64
	windows-386
	openbsd-amd64
	openbsd-386
	freebsd-amd64
	freebsd-386
	netbsd-amd64
	netbsd-386
)

for platform in "${platforms[@]}"; do
	goos=${platform%-*}
	goarch=${platform##*-}
	ext=""
	if [ "$goos" = "windows" ]; then
		ext=".exe"
	fi
	GOOS=${platform%-*} GOARCH=${platform##*-} go build -o "dist/git-lzc-$platform$ext" ./cmd/git-lzc
done


