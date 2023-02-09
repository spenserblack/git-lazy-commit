package lazycommit

import "testing"

// Tests that a commit message can't be built when there are no staged changes.
func TestBuildCommitMessageNoStaged(t *testing.T) {
	dir := tempRepo(t)
	repo, err := OpenRepo(dir)
	if err != nil {
		t.Fatal(err)
	}
	_, err = repo.CommitMsg()
	if err == nil {
		t.Fatal("expected error")
	}
}
