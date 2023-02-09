package lazycommit

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5"

	"github.com/spenserblack/git-lazy-commit/pkg/fileutils"
)

// CommitMsg builds a commit message using the tracked files in the repository.
func (r *LazyRepo) CommitMsg() (string, error) {
	status, err := r.status()
	if err != nil {
		return "", err
	}

	if len(status) == 0 {
		return "", errors.New("no tracked files")
	}
	if len(status) == 1 {
		for filename, fileStatus := range status {
			return singleFileMsg(filename, fileStatus), nil
		}
	}
	return multiFileMsg(status), nil
}

func singleFileMsg(filename string, fileStatus *git.FileStatus) string {
	statusString := ""
	switch fileStatus.Staging {
	case git.Added:
		statusString = "Create"
	case git.Deleted:
		statusString = "Delete"
	case git.Modified:
		statusString = "Update"
	case git.Renamed:
		statusString = "Rename to"
	case git.Copied:
		statusString = "Copy to"
	default:
		statusString = "Do something to"
	}

	return fmt.Sprintf("%s %s", statusString, filename)
}

func multiFileMsg(status git.Status) string {
	var builder strings.Builder

	filenames := make([]string, 0, len(status))
	for name := range status {
		filenames = append(filenames, name)
	}

	sharedDir := fileutils.SharedDirectory(filenames)

	if sharedDir == "/" {
		builder.WriteString("Update files\n")
	} else {
		builder.WriteString(fmt.Sprintf("Update %s/\n", sharedDir))
	}
	builder.WriteRune('\n')

	for filename, fileStatus := range status {
		msgItem := singleFileMsg(filename, fileStatus)
		builder.WriteString(fmt.Sprintf("- %s\n", msgItem))
	}

	return builder.String()
}
