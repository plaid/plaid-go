# Publish guide

`plaid-go` is published on this repo via tags, as it uses GOMODULES.

After merging into our [beta branch](https://github.com/plaid/plaid-go/tree/0.1.0-beta-release), run the following commands.

```
git tag v{VERSION}
git push origin v{VERSION}
```

Make sure you have the `v` prefix! Also, we are denoting beta versions with both sub `1.0` versions plus with the suffix `-beta.x`. e.g. `v0.1.0-beta.1`.

After you do this, you've released a new version of `plaid-go`!
