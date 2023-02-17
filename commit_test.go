package lazycommit

import (
	"strings"
	"testing"
)

// Tests that a commit message can't be built when there are no staged changes.
func TestBuildCommitMessage(t *testing.T) {
	t.Log("Creating a new repo.")
	dir := tempRepo(t)
	repo := Repo(dir)

	_, err := repo.CommitMsg()
	if err == nil && err.Error() != "no tracked files" {
		t.Errorf(`Expected "no tracked files", got %v`, err)
	}

	f := commitFile(t, dir, "test.txt", "test")
	defer f.Close()

	t.Log(`Modifying test.txt`)
	commitFile(t, dir, "test.txt", "")
	addFile(t, dir, "test.txt", "different text")

	msg, err := repo.CommitMsg()
	if err != nil {
		t.Fatal(err)
	}
	if msg != "Update test.txt" {
		t.Errorf(`Expected "Update test.txt", got %v`, msg)
	}

	t.Log(`Adding a new file`)
	addFile(t, dir, "test2.txt", "test")

	msg, err = repo.CommitMsg()
	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(msg, "\n")
	if lines[0] != "Update files" {
		t.Errorf(`Expected "Update files" in the header, got %v`, lines[0])
	}
	if lines[1] != "" {
		t.Errorf(`Expected an empty line after the header, got %v`, lines[1])
	}
	body := strings.Join(lines[2:], "\n")
	t.Logf("Body:\n %v", body)
	for _, want := range []string{"- Update test.txt", "- Create test2.txt"} {
		if !strings.Contains(body, want) {
			t.Errorf(`Expected %v in the body`, want)
		}
	}
}

// TestBuildCommitMessageWithRename tests that a commit message can be built when a file is renamed.
func TestBuildCommitMessageWithRename(t *testing.T) {
	dir := tempRepo(t)
	repo := Repo(dir)

	f := commitFile(t, dir, "foo.txt", "test")
	defer f.Close()

	t.Log(`Renaming test.txt to test2.txt`)
	moveFile(t, dir, "foo.txt", "bar.txt")

	msg, err := repo.CommitMsg()
	if err != nil {
		t.Fatal(err)
	}
	if msg != "Rename foo.txt to bar.txt" {
		t.Errorf(`Expected "Rename foo.txt to bar.txt", got %v`, msg)
	}
}
