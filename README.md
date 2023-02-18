# Lazy Commit

[![CI](https://github.com/spenserblack/git-lazy-commit/actions/workflows/ci.yml/badge.svg)](https://github.com/spenserblack/git-lazy-commit/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/spenserblack/git-lazy-commit/branch/main/graph/badge.svg?token=nFiCRNnexU)](https://codecov.io/gh/spenserblack/git-lazy-commit)
[![Go Report Card](https://goreportcard.com/badge/github.com/spenserblack/git-lazy-commit)](https://goreportcard.com/report/github.com/spenserblack/git-lazy-commit)

Be lazy and just commit

## Description

This provides the `git lzc` command. This command can be used
in situations where you don't really care about choosing which
changes to track or writing your commit message -- you just want to
commit your work.

By its nature, `git lzc` can very easily add accidental changes
to the git history if the user isn't careful. So, while this
tool may be appealing to git beginners, its target audience is
actually experienced git users who know when they want to break
the rules for creating good commits.

### Staged Changes

If you have staged changes (`git add path/to/file`), then
`git lzc` will commit those staged changes. If you *do not*
have any staged changes, then `git lzc` will commit *all* changes,
**including untracked files** (so be careful!).

### Commit Messages

`git lzc` will write your commit message for you. If you've changed
a single file, the commit message will look like this:

```
Update www/index.html
```

If you've changed multiple files that share a similar directory, your
commit message will look like this:

```
Update public/

- Update public/favicon.ico
- Create public/icons/favicon-16x16.png
- Create public/icons/favicon-32x32.png
```

If there aren't any similar directories that all changes share, or at least one
of the updated files is in the root of the repository, your commit message
will look like this:

```
Update files

- Update views.py
- Update templates/myapp/index.html
```

## Installation

### Unix

```shell
curl https://raw.githubusercontent.com/spenserblack/git-lazy-commit/main/install.sh | sh
```

## Windows (PowerShell)

You may need to run this as an administrator.

```powershell
Invoke-WebRequest "https://raw.githubusercontent.com/spenserblack/git-lazy-commit/main/install.ps1" | Invoke-Expression
```

### From GitHub Releases

Download the appropriate executable from the [release assets][latest-release],
rename it to `git-lzc`, and add it to a location in `PATH`.

[latest-release]: https://github.com/spenserblack/git-lazy-commit/releases/latest
