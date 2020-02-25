# Publishing

`github.plaid.com/plaid/plaid-go` uses semantic versioning with git tags to
track releases.

1. You will need write access to master. If you do not have it, request it from a member of the Plaid team.

2. Ensure that you have cloned this repository into your `GOPATH` and turned `GO111MODULE=off`

3. On the master branch, use `make release-(patch|minor|major)` to prepare the
repository for release. Verify the changes made to the git repository, and
then run `git push --follow-tags` to publish the release to github.
