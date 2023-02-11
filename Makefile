FILES = $(shell find -type f -name '*.go')

git-lzc: $(FILES)
	go build ./cmd/git-lzc/
