package main

import (
	"fmt"

	"github.com/spf13/cobra"

	lazycommit "github.com/spenserblack/git-lazy-commit"
)

var rootCmd = &cobra.Command{
	Use:   "git-lzc",
	Short: "Lazy Commit generates commit messages for you",
	Long: `Lazy Commit checks your git status, stages files if they're all unstaged,
				and generates a commit message for you.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		repo := lazycommit.Repo(".")

		noStaged, err := repo.NoStaged()
		onError(err)

		if noStaged {
			onError(repo.StageAll())
		}

		out, err := repo.Commit()
		onError(err)

		fmt.Printf("%s", out)
	},
	SilenceUsage: true,
}
