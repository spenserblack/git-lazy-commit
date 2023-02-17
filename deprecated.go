
package lazycommit

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
	return r.wt.AddWithOptions(&git.AddOptions{All: true})
}

// Status gets the repo's status.
func (r *LazyRepo) status() (git.Status, error) {
	return r.wt.Status()
}
