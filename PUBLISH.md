# Publishing

`github.plaid.com/plaid/plaid-go` uses semantic versioning with git tags to
track releases.

On the master branch, use `make release-(patch|minor|major)` to prepare the
repository for release. Verify the changes made to the git repository, and
then run `git push --follow-tags` to publish the release to github.
