package main

import (
	"fmt"
	"os"
	"strings"

	lazycommit "github.com/spenserblack/git-lazy-commit"
)

func main() {
	repo, err := lazycommit.OpenRepo(".")
	onError(err)

	noStaged, err := repo.NoStaged()
	onError(err)

	if noStaged {
		onError(repo.StageAll())
	}

	hash, msg, err := repo.Commit()
	onError(err)

	msgLines := strings.Split(msg, "\n")

	fmt.Printf("[%s] %s\n", hash, msgLines[0])
}

func onError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
