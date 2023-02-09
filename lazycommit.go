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

// NoStaged checks if there are no staged changes (added files, changed files, removed files)
// in the repository.
func (r *LazyRepo) NoStaged() (bool, error) {
	status, err := r.status()
	if err != nil {
		return false, err
	}

	for _, file := range status {
		if file.Staging != git.Unmodified {
			return false, nil
		}
	}

	return true, nil
}

// StageAll stages all changes in the repository.
func (r *LazyRepo) StageAll() error {
	wt, err := (*git.Repository)(r).Worktree()
	if err != nil {
		return err
	}
	return wt.AddWithOptions(&git.AddOptions{All: true})
}

// Worktree gets the repo's worktree.
func (r *LazyRepo) worktree() (*git.Worktree, error) {
	return (*git.Repository)(r).Worktree()
}

// Status gets the repo's status.
func (r *LazyRepo) status() (git.Status, error) {
	wt, err := r.worktree()
	if err != nil {
		return nil, err
	}
	return wt.Status()
}
