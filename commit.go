package lazycommit

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spenserblack/git-lazy-commit/pkg/fileutils"
)

// CommitMsg builds a commit message using the tracked files in the repository.
func (repo Repo) CommitMsg() (string, error) {
	statuses, err := repo.Status()
	if err != nil {
		return "", err
	}

	// NOTE: Filtering to only statuses that are staged and can be used for the commit message.
	commitableStatuses := make([]StatusRecord, 0, len(statuses))
	for _, status := range statuses {
		if _, ok := statusMap[status.Staged]; ok {
			commitableStatuses = append(commitableStatuses, status)
		}
	}

	if len(commitableStatuses) == 0 {
		return "", errors.New("no tracked files")
	}

	if len(commitableStatuses) == 1 {
		status := commitableStatuses[0]
		return fmt.Sprintf("%s %s", statusMap[status.Staged], status.Path), nil
	}

	return multiFileMsg(commitableStatuses), nil
}

// MultiFileMsg builds a commit message from multiple files.
func multiFileMsg(statuses []StatusRecord) string {
	var builder strings.Builder
	filenames := make([]string, 0, len(statuses))
	for _, status := range statuses {
		filenames = append(filenames, status.Path)
	}

	sharedDir := fileutils.SharedDirectory(filenames)

	if sharedDir == "/" {
		builder.WriteString("Update files\n")
	} else {
		builder.WriteString(fmt.Sprintf("Update %s/\n", sharedDir))
	}
	builder.WriteRune('\n')

	for _, status := range statuses {
		builder.WriteString(fmt.Sprintf("- %s %s\n", statusMap[status.Staged], status.Path))
	}

	return builder.String()
}
