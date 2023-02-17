package lazycommit

import (
	"strings"
)

// StatusMap maps status codes from "git status --porcelain" to human-readable, imperative
// verbs.
var statusMap = map[rune]string{
	'M': "Update",
	'A': "Create",
	'D': "Delete",
	// NOTE: With -z, the *new* filename is followed by the old filename, separated by a NUL.
	'R': "Rename",
	'C': "Copy",
	// NOTE: '?' is untracked, ' ' is unmodified
}

// Status runs "git status --porcelain -z" parses the output and returns a list of
// files with their status.

// NoStaged checks if there are no staged changes (added files, changed files, removed files)
// in the repository.
func (repo Repo) NoStaged() (bool, error) {
	cmd, err := repo.cmd("status", "--porcelain", "-z", "--untracked-files=no")
	if err != nil {
		return false, err
	}
	out, err := cmd.Output()
	if err != nil {
		return false, err
	}
	statuses := strings.Split(string(out), "\x00")
	for _, status := range statuses {
		if len(status) == 0 {
			continue
		}
		stagedStatus := status[0]
		if stagedStatus != ' ' && stagedStatus != '?' {
			return false, nil
		}
	}
	return true, nil
}
