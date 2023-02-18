FILES = $(shell find -type f -name '*.go')
DISTRIBUTIONS = darwin-amd64 linux-amd64 windows-amd64 openbsd-amd64 freebsd-amd64 netbsd-amd64 linux-386 windows-386 openbsd-386 freebsd-386 netbsd-386
targets = $(foreach distribution,$(DISTRIBUTIONS),dist/git-lzc-$(distribution)$(if $(findstring windows,$(distribution)),.exe))

git-lzc: $(FILES) go.mod go.sum
	go build ./cmd/git-lzc/

$(targets): dist/git-lzc-%: $(FILES) go.mod go.sum
	GOOS=$(word 1, $(subst -, ,$*)) GOARCH=$(word 1, $(subst ., ,$(word 2, $(subst -, ,$*)))) go build -o $@ ./cmd/git-lzc/

.PHONY: all
all: $(targets)

.PHONY: install
install: git-lzc
	cp git-lzc /usr/local/bin/
