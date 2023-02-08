package lazycommit

import "testing"

// Tests that OpenRepo returns a LazyRepo if the repository can be opened.
func TestOpenRepo(t *testing.T) {
	dir := tempRepo(t)
	var (
		repo *LazyRepo
		err  error
	)
	repo, err = OpenRepo(dir)
	if err != nil {
		t.Fatal(err)
	}
	if repo == nil {
		t.Fatal("repo is nil")
	}
}

// Tests that, if a repo cannot be opened, OpenRepo returns an error.
func TestOpenRepoError(t *testing.T) {
	repo, err := OpenRepo(t.TempDir())
	if err == nil {
		t.Fatal("expected error")
	}
	if repo != nil {
		t.Fatal("expected repo to be nil")
	}
}
