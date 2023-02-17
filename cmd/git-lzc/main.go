package main

import (
	"fmt"
	"os"

	lazycommit "github.com/spenserblack/git-lazy-commit"
)

func main() {
	repo := lazycommit.Repo(".")

	noStaged, err := repo.NoStaged()
	onError(err)

	if noStaged {
		onError(repo.StageAll())
	}

	out, err := repo.Commit()
	onError(err)

	fmt.Printf("%s", out)
}

func onError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
