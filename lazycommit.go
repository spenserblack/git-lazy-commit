// Package lazycommit mostly provides wrappers around go-git to make it easier for
// "lazy" usage.
package lazycommit

import "github.com/go-git/go-git/v5"

// LazyRepo is a wrapper around go-git's Repository for simpler usage.
type LazyRepo git.Repository

// OpenRepo opens a repository at the given path.
func OpenRepo(path string) (*LazyRepo, error) {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}
	return (*LazyRepo)(repo), nil
}
