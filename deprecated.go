package lazycommit

import (
	"github.com/go-git/go-git/v5/plumbing"
)

// LazyRepo is a wrapper around go-git's Repository for simpler usage.
//
// DEPRECATED: Use Repo instead.
type LazyRepo struct {
	r Repo
}

// OpenRepo returns a LazyRepo for compatibility with previous versions of lazycommit.
//
// DEPRECATED: See note on LazyRepo.
func OpenRepo(path string) (*LazyRepo, error) {
	return &LazyRepo{Repo(path)}, nil
}

// NoStaged checks if there are no staged changes (added files, changed files, removed files)
// in the repository.
func (r *LazyRepo) NoStaged() (bool, error) {
	return r.r.NoStaged()
}

// StageAll stages all changes in the repository.
func (r *LazyRepo) StageAll() error {
	return r.r.StageAll()
}

// Commit commits all changes in the repository.
//
// It returns the commit hash and the commit message.
func (r *LazyRepo) Commit() (hash plumbing.Hash, msg string, err error) {
	panic("Incompatible types")
}

// CommitMsg builds a commit message using the tracked files in the repository.
func (r *LazyRepo) CommitMsg() (string, error) {
	return r.r.CommitMsg()
}
