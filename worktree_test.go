package lazycommit

import (
	"os/exec"
	"testing"
)

// TestStageAll tests that all changes are staged.
func TestStageAll(t *testing.T) {
	dir := tempRepo(t)
	repo := Repo(dir)

	writeFile(t, dir, "test.txt", "test")
	writeFile(t, dir, "test2.txt", "test")

	err := repo.StageAll()
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		filename string
		want    string
	}{
		{"test.txt", "A  test.txt\x00"},
		{"test2.txt", "A  test2.txt\x00"},
	}

	for _, tt := range tests {
		cmd := exec.Command("git", "status", "--porcelain", "-z", "--", tt.filename)
		cmd.Dir = dir
		out, err := cmd.Output()
		if err != nil {
			t.Fatal(err)
		}
		if string(out) != tt.want {
			t.Errorf("expected %s status to be %q, got %q", tt.filename, tt.want, string(out))
		}
	}
}
