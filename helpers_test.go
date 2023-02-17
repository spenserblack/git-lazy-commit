package lazycommit

import (
	"os"
	"os/exec"
	"path"
	"testing"
)

// Helper function to create a new repository in a temporary directory.
// Returns the path to the repository.
func tempRepo(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	cmd := exec.Command("git", "init")
	cmd.Dir = dir
	err := cmd.Run()
	if err != nil {
		t.Fatal(err)
	}
	return dir
}

// Helper function that writes a file, but does not stage or commit it.
func writeFile(t *testing.T, dir, filename, contents string) *os.File {
	t.Helper()
	f, err := os.Create(path.Join(dir, filename))
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.WriteString(contents)
	if err != nil {
		t.Fatal(err)
	}
	return f
}

// Helper function that writes a file and stages it (but doesn't commit it).
func addFile(t *testing.T, dir, filename, contents string) *os.File {
	t.Helper()
	f := writeFile(t, dir, filename, contents)
	cmd := exec.Command("git", "add", filename)
	cmd.Dir = dir
	err := cmd.Run()
	if err != nil {
		t.Fatal(err)
	}
	return f
}

// Helper function that commits a file to the repository.
func commitFile(t *testing.T, dir, filename, contents string) *os.File {
	t.Helper()
	f := addFile(t, dir, filename, contents)
	cmd := exec.Command("git", "commit", "-m", "test")
	cmd.Dir = dir
	err := cmd.Run()
	if err != nil {
		t.Fatal(err)
	}
	return f
}

// Helper function that moves a file.
func moveFile(t *testing.T, dir, oldName, newName string) {
	t.Helper()
	cmd := exec.Command("git", "mv", oldName, newName)
	cmd.Dir = dir
	err := cmd.Run()
	if err != nil {
		t.Fatal(err)
	}
}
