package lazycommit

import (
	"testing"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-git/v5"
	gitconfig "github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// Helper function to create a new repository in a temporary directory.
// Returns the path to the repository.
func tempRepo(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	_, err := git.PlainInit(dir, false)
	if err != nil {
		t.Fatal(err)
	}
	return dir
}

// Helper function that writes a file, but does not stage or commit it.
func writeFile(t *testing.T, dir, filename, contents string) (*git.Worktree, billy.File) {
	t.Helper()
	rawRepo, err := git.PlainOpen(dir)
	if err != nil {
		t.Fatal(err)
	}
	wt, err := rawRepo.Worktree()
	if err != nil {
		t.Fatal(err)
	}
	f, err := wt.Filesystem.Create(filename)
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.Write([]byte(contents))
	if err != nil {
		t.Fatal(err)
	}

	return wt, f
}

// Helper function that writes a file and stages it (but doesn't commit it).
func addFile(t *testing.T, dir, filename, contents string) billy.File {
	t.Helper()
	wt, f := writeFile(t, dir, filename, contents)
	_, err := wt.Add(filename)
	if err != nil {
		t.Fatal(err)
	}

	return f
}

// Helper function that commits a file to the repository.
func commitFile(t *testing.T, dir, filename, contents string) billy.File {
	t.Helper()
	rawRepo, err := git.PlainOpen(dir)
	if err != nil {
		t.Fatal(err)
	}
	wt, err := rawRepo.Worktree()
	if err != nil {
		t.Fatal(err)
	}
	f := addFile(t, dir, filename, contents)
	_, err = wt.Commit("test commit", &git.CommitOptions{
		AllowEmptyCommits: true,
		Author: &object.Signature{
			Name:  "Test",
			Email: "test@example.com",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	return f
}

// Helper function that gets the working tree of a repository.
func getWorktree(t *testing.T, dir string) *git.Worktree {
	t.Helper()
	rawRepo, err := git.PlainOpen(dir)
	if err != nil {
		t.Fatal(err)
	}
	wt, err := rawRepo.Worktree()
	if err != nil {
		t.Fatal(err)
	}
	return wt
}

// Helper function that gets the status of a repository.
func getStatus(t *testing.T, dir string) git.Status {
	t.Helper()
	wt := getWorktree(t, dir)
	status, err := wt.Status()
	if err != nil {
		t.Fatal(err)
	}
	return status
}

// Helper function that updates a repo's config.
func updateConfig(t *testing.T, dir string, f func(*gitconfig.Config)) {
	t.Helper()
	rawRepo, err := git.PlainOpen(dir)
	if err != nil {
		t.Fatal(err)
	}
	config, err := rawRepo.Config()
	if err != nil {
		t.Fatal(err)
	}
	f(config)
	err = rawRepo.Storer.SetConfig(config)
	if err != nil {
		t.Fatal(err)
	}
}
