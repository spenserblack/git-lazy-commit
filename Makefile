FILES = $(shell find -type f -name '*.go')

git-lzc: $(FILES) go.mod go.sum
	go build ./cmd/git-lzc/

.PHONY: install
install: git-lzc
	cp git-lzc /usr/local/bin/
