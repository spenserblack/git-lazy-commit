package lazycommit

import (
	"testing"

	"github.com/go-git/go-git/v5"
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

