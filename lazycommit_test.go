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

// Tests that NoStaged returns true if there are no staged changes.
func TestNoStaged(t *testing.T) {
	dir := tempRepo(t)
	// NOTE: Committing a file so that there's something in the worktree.
	f := commitFile(t, dir, "test.txt", "test")
	// NOTE: Adding some unstaged contents to the file
	_, err := f.Write([]byte("changes"))
	if err != nil {
		t.Fatal(err)
	}

	repo, err := OpenRepo(dir)
	if err != nil {
		t.Fatal(err)
	}
	noStaged, err := repo.NoStaged()
	if err != nil {
		t.Fatal(err)
	}
	if !noStaged {
		t.Fatal("expected no staged changes")
	}
}

// Tests that NoStaged returns false if there are staged changes.
func TestNoStagedStaged(t *testing.T) {
	dir := tempRepo(t)
	// NOTE: Committing a file so that there's something in the worktree.
	commitFile(t, dir, "test.txt", "test")

	repo, err := OpenRepo(dir)
	if err != nil {
		t.Fatal(err)
	}
	addFile(t, dir, "test2.txt", "test")

	noStaged, err := repo.NoStaged()
	if err != nil {
		t.Fatal(err)
	}
	if noStaged {
		t.Fatal("expected staged changes")
	}
}
