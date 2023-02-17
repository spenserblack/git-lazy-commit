package lazycommit

import "testing"

// TestNoStagedChanges tests that NoStaged returns true if there are no staged changes, and false otherwise.
func TestNoStagedChanges(t *testing.T) {
	t.Log("Creating a new repo.")
	dir := tempRepo(t)
	repo := Repo(dir)
	noStaged, err := repo.NoStaged()
	if err != nil {
		t.Fatal(err)
	}
	if !noStaged {
		t.Error("expected no staged changes")
	}

	t.Log("Committing a file so that there's something in the worktree.")
	f := commitFile(t, dir, "test.txt", "test")
	defer f.Close()

	noStaged, err = repo.NoStaged()
	if err != nil {
		t.Fatal(err)
	}
	if !noStaged {
		t.Error("expected no staged changes")
	}

	t.Log("Adding a staged file.")
	addFile(t, dir, "test2.txt", "test")

	noStaged, err = repo.NoStaged()
	if err != nil {
		t.Fatal(err)
	}

	if noStaged {
		t.Error("expected staged changes")
	}
}
