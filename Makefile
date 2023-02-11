FILES = $(shell find -type f -name '*.go')

git-lzc: $(FILES) go.mod go.sum
	go build ./cmd/git-lzc/
