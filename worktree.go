package lazycommit

// StageAll stages all changes in the repository.
func (repo Repo) StageAll() error {
	cmd, err := repo.cmd("add", "--all")
	if err != nil {
		return err
	}
	return cmd.Run()
}
