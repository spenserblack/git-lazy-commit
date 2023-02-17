package lazycommit

import (
	"testing"

	"github.com/go-git/go-git/v5"
)

// Tests that StageAll stages all changes in the repository.
func TestStageAll(t *testing.T) {
	dir := tempRepo(t)
	// NOTE: Committing a file so that there's something in the worktree.
	commitFile(t, dir, "test.txt", "test")
	writeFile(t, dir, "test2.txt", "test")

	repo, err := OpenRepo(dir)
	if err != nil {
		t.Fatal(err)
	}
	err = repo.StageAll()
	if err != nil {
		t.Fatal(err)
	}

	status := getStatus(t, dir)

	if fileStatus := status.File("test2.txt"); fileStatus.Staging != git.Added {
		t.Errorf("expected test2.txt to be staged, got %v", fileStatus.Staging)
	}
}
