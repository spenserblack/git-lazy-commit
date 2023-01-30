# Lazy Commit

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
Update index.html

- Update www/index.html
```

If you've changed multiple files that share a top-level directory, your
commit message will look like this:

```
Update public/

- Update public/favicon.ico
- Create public/icons/favicon-16x16.png
- Create public/icons/favicon-32x32.png
```

If your changes span multiple top-level directories, or at least one
of the updated files is in the root of the repository, your commit message
will look like this:

```
Update files

- Update views.py
- Update templates/myapp/index.html
```
