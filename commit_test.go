package lazycommit

import (
	"strings"
	"testing"

	gitconfig "github.com/go-git/go-git/v5/config"
)

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

// Tests that commit commits all files in the worktree.
func TestCommit(t *testing.T) {
	dir := tempRepo(t)
	updateConfig(t, dir, func(config *gitconfig.Config) {
		config.User.Name = "Test User"
		config.User.Email = "test@example.com"
	})
	addFile(t, dir, "test.txt", "test")
	addFile(t, dir, "test2.txt", "test")

	repo, err := OpenRepo(dir)
	if err != nil {
		t.Fatal(err)
	}

	_, msg, err := repo.Commit()
	if err != nil {
		t.Fatal(err)
	}

	wantHeader := "Update files"
	wantBodyLines := []string{"- Create test.txt", "- Create test2.txt"}

	if !strings.HasPrefix(msg, wantHeader) {
		t.Errorf("expected commit message to start with %q, got %q", wantHeader, msg)
	}

	for _, line := range wantBodyLines {
		if !strings.Contains(msg, line) {
			t.Errorf("expected commit message to contain %q, got %q", line, msg)
		}
	}
}
