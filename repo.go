package lazycommit

import (
	"os/exec"

	"github.com/cli/safeexec"
)

// Repo is a path to a git repository. Used to call the "git" command on the
// repository.
type Repo string

// Cmd runs a git command on the repository.
func (repo Repo) cmd(args ...string) (*exec.Cmd, error) {
	gitBin, err := safeexec.LookPath("git")
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(gitBin, args...)
	cmd.Dir = string(repo)
	return cmd, nil
}
