package lazycommit

import "os/exec"

// Repo is a path to a git repository. Used to call the "git" command on the
// repository.
type Repo string

// Cmd runs a git command on the repository.
func (repo Repo) cmd(args ...string) *exec.Cmd {
	cmd := exec.Command("git", args...)
	cmd.Dir = string(repo)
	return cmd
}
