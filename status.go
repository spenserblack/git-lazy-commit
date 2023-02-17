package lazycommit

import (
	"strings"
)

// StatusRecord represents a single status record from "git status".
type StatusRecord struct {
	// Staged is the staged status of the file.
	Staged rune
	// Unstaged is the unstaged status of the file.
	Unstaged rune
	// Path is the path to the file.
	Path string
	// Dest is the destination path for a rename or copy.
	Dest string
}

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

// NoStaged checks if there are no staged changes (added files, changed files, removed files)
// in the repository.
func (repo Repo) NoStaged() (bool, error) {
	statuses, err := repo.Status()
	if err != nil {
		return false, err
	}
	for _, status := range statuses {
		if status.Staged != ' ' && status.Staged != '?' {
			return false, nil
		}
	}
	return true, nil
}

// Status gets and parses the repo's status.
func (repo Repo) Status() ([]StatusRecord, error) {
	// TODO: Test this method with a variety of added, moved, deleted, and modified files.
	cmd, err := repo.cmd("status", "--porcelain", "-z")
	if err != nil {
		return nil, err
	}
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	statuses := strings.Split(string(out), "\x00")
	records := make([]StatusRecord, 0, len(statuses))

	for i := 0; i < len(statuses); i++ {
		status := []rune(statuses[i])
		if len(status) == 0 {
			continue
		}
		stagedStatus := status[0]
		unstagedStatus := status[1]
		path := string(status[3:])
		dest := ""
		if unstagedStatus == 'R' || unstagedStatus == 'C' {
			dest = path
			i++
			path = statuses[i]
		}
		records = append(records, StatusRecord{
			Staged:   stagedStatus,
			Unstaged: unstagedStatus,
			Path:     path,
			Dest:     dest,
		})
	}

	return records, nil
}
