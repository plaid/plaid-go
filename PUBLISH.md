# Publish guide

`plaid-go` is published on this repo via tags, as it uses GOMODULES. Follow SEMVER to come up with the new `VERSION` number. Make sure to first update the `CHANGELOG` file.

After merging into `master`, run the following commands.

```
git tag v{VERSION}
git push origin v{VERSION}
```

Make sure you have the `v` prefix!

After you do this, you've released a new version of `plaid-go`!
