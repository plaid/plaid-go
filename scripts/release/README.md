# Automated Go SDK publication

After a reviewed change merges to `plaid/plaid-openapi`, generation produces both
Go clients and a signed record of their exact inputs and content. The integration
workflow checks the generated PR and required CI, then merges it automatically.
The release workflow verifies that merge and generation record, builds and tests
the SDK, attests a canonical Go module ZIP, and publishes a signed version tag.
The generation pipeline selects a minor release for compatible exported Go APIs
and the next major version when either client's exported API changes incompatibly.
No additional SDK PR approval or workflow dispatch is required. Automatic
integration and publication remain disabled until configured and tested.

A separately human-reviewed SDK PR remains an explicit manual fallback. Neither
path may ignore an outstanding changes-requested review or alter an existing tag.

The SDK and `legacy/` must come from the same public OpenAPI spec. The generation
record binds the spec, public config commit, immutable generator images, version,
legacy tree and all tracked SDK content. Public Fern metadata contains only the
CLI, generator and SDK versions; the signed manifest records the shared spec/config commit.

## Major-version transitions

The source generation workflow selects and signs the version. The integration
policy permits the existing module major or exactly the next major at `.0.0`;
it rejects skipped majors, downgrades and partial or prerelease major transitions.
Both clients must remain bundled under the new root module.

During a major transition, only the exact old `github.com/plaid/plaid-go/vN`
module token may change to the next major in the explicitly maintained Go files
and README. This includes imports, comments and examples, but no other edits.
Similarly named module paths are untouched. Workflows, release scripts, CI,
license, ignore files and file modes remain unchanged. Same-major releases retain
the previous byte-for-byte preservation checks.

The publisher independently checks the version against existing tags and keeps
the exact-manual-predecessor and public-spec ancestry requirements across majors.
A new major still requires signed generation evidence, trusted CI, and the same
server-side protections. API comparison cannot prove runtime compatibility;
generator errors and failing checks remain failures, not reasons to publish.

## Approved generated paths

Generated PR changes are limited to lowercase Go source paths and these exact
resources: `go.mod`, `go.sum`, `.fern/metadata.json`, `reference.md`, and
`wiremock/wiremock-mappings.json`. Generated files must use non-executable mode
`100644`; new endpoint package directories need no policy update. Changes outside
that list stop integration, including added scripts/configuration and deletion of
existing unlisted files.

The maintained-file rules take precedence over the Go allowlist. Everything under
`legacy/tests/` is maintained, with only the exact major module-path substitution
allowed in existing Go files. New or removed legacy tests require separate review.
README, changelog and release evidence retain their narrower existing checks.
A signed generation record establishes origin and exact content; it does not
expand these approved paths or establish that generated code is correct.

## Separation of permissions

- Verification reads public source and its public generation signature. The
  spec and config must come from the exact same `plaid/plaid-openapi` commit.
  No private repository or private-attestation credential is needed.
- SDK execution runs in a job without publishing credentials, with an allowlisted
  environment. Packaging uses a fresh checkout in a separate job, so tests cannot
  change what is packaged.
- Attestation verifies the package and signs its ZIP and release manifest. The
  publication preflight runs without tag-writing credentials, including when a
  manually requested dry run explicitly permits a test attestation.
- Publication uses the reviewed control scripts from an immutable commit and a
  bare Git repository. It rechecks the protected branch, authorized merge, version floor,
  spec ancestry, artifact signature and remote content before signing and pushing
  only the new tag. It does not execute SDK code. An existing correctly signed tag
  at the exact commit is an idempotent success; a conflicting tag is an error.

## Required GitHub setup

Verify that GitHub accepts the CODEOWNERS entries. Keep strict, up-to-date required
SDK CI in classic branch protection on `master`, with no classic PR approval
requirement or classic approval bypass. Human approval is enforced separately by
an active repository branch ruleset that explicitly includes `refs/heads/master`
without exclusions. It must contain only a pull-request rule requiring at least
one approval, code-owner review, and dismissal of stale reviews.

Set `SDK_REVIEW_RULESET_ID` to that ruleset's numeric ID. Give only the integration
App bypass on this approval ruleset with mode `pull_request` ("For pull requests
only"). Do not use `always` or `exempt`, and do not place required CI in the
bypassed ruleset. PR-only mode prevents direct pushes; it applies to the App's PR
operations, not to specific generated files. Our verifier enforces the generated
content restrictions and checks the App's effective PR-only capability using the
exact installation token about to merge. The read-only phase checks ruleset
identity, applicability and review requirements without assuming its token can merge.

GitHub can omit the full bypass list from API responses to limited tokens. Runtime
verification checks our merger's effective capability; the setup audit must still
confirm that no other identity has unintended bypass rights. A rulesets-only CI
setup remains unsupported: the verifier deliberately requires classic strict CI
alongside the native PR-only approval ruleset. Hosted tests must demonstrate that
passing direct pushes fail, passing generated PRs merge, and required CI cannot
be bypassed. A bot must never post a rubber-stamp approval to satisfy review rules.

Require `sdk-generation/verified` in the same classic strict status checks, bound
to the GitHub Actions App (integration ID `15368`), alongside CircleCI. This is a
server-enforced condition: the merger App has no status-write permission and
cannot satisfy it by calling the merge API directly.

Only the trusted integration check job receives `statuses: write`. It reports a
status only when the verified outcome changes, avoiding GitHub's per-commit status
limit. Success requires complete signed-generation and
approved-path checks, or after exact-head human approval and GitHub's `APPROVED`
review decision for a manual PR. Pending checks or missing reviews leave pending;
verification errors mark failure. This preserves separately reviewed SDK/control
changes without granting an author-based exemption. The manual path does not
create generation evidence or make an unsigned candidate publishable.

The verifier ignores its own pending status while checking other CI, avoiding a
circular dependency; the merge boundary and GitHub both require its success.
A zero-permission review signal workflow triggers a trusted default-branch rescan
when reviews are submitted, dismissed or edited. No signal-run code or artifacts
are executed by the trusted workflow, and its own status updates do not retrigger
it. Scheduled checks remain a fallback.

Merge these controls and the review-signal workflow through the existing reviewed
path first, then update the immutable control pin. Configure the approval ruleset
and both classic required statuses before setting `SDK_REVIEW_RULESET_ID`. A
configured ruleset ID and control pin enable status verification independently of
`SDK_AUTOMERGE_ENABLED`; manual PRs can receive verification while automatic
merging stays off. Manual dry runs still write no statuses. Use a normal dispatch
with automatic merging disabled to exercise status reporting without merging.

A control-workflow PR is checked against the current pinned base, so it can obtain
human-review verification before its new workflow lands. Refresh the control pin
after merging it. Turning automatic merging off retains verification for manual
PRs. If retiring status verification itself, first remove its required status
through the reviewed repository-configuration process; then remove credentials
or settings. Otherwise manual PRs would be left waiting for a status no job emits.

App binding identifies GitHub Actions, not an individual workflow. Other workflows
must not give status-write credentials to candidate code. Reviews, CI and statuses
can change between checks and merging; notifications reduce stale-status windows
but are not an atomic per-PR approval exception. The normal merger still rechecks
live CI, exact head/base and human objections immediately before its merge call.

Create `sdk-integration` and `release` environments restricted to protected
`master`, with no per-run human approval on the automatic path. The integration
App has Contents write, Pull requests read, Commit statuses read and Administration
read (for checking branch protection), scoped only to this SDK repository. Its check job obtains a
read-only token for inspections, but that job still holds the full App private
key and must be trusted with its full authority. Generated SDK code never runs
there. Only the merge command is passed a write-capable installation token.
Neither App gets Administration write or workflow write.

Contents write permits Git ref operations; role names do not enforce restrictions.
Configure a tag-creation rule allowing the publisher and approved manual-release
actors, plus a separate rule blocking tag updates/deletions with no automation
bypass. Creator and merger identities must not bypass tag creation rules. The
publisher must not bypass branch review requirements. Keep required CI outside
any review exception. Prove these restrictions with the actual identities before
activation; the scripts alone cannot enforce what a stolen credential can do.

The SDK PR creator is a third, separate identity with no review bypass. Pin the
numeric bot user IDs of the creator and integrator; display names/labels do not
authorize a merge or release. Generated PRs cannot modify the trusted workflows,
release scripts, CI, license or maintained customizations. Empty/missing identity
settings stop automatic integration. Set `SDK_FAILURE_ASSIGNEE` to a repository
maintainer. A separate failure workflow creates or updates one assigned issue per
failed integration/release workflow, linking the failed run without copying logs.
It ignores forks, non-master runs and PR-triggered jobs, and has only Issues write.
Confirm the assignee actually receives notifications before activation; issue
assignment does not guarantee email or Slack delivery. Close the issue after
recovery. Pending CI is a normal wait; failed CI produces an integration failure.
Repository variables:

| Variable | Value |
|---|---|
| `SDK_FAILURE_ASSIGNEE` | Assignable maintainer who handles failed automation runs |
| `SDK_AUTOMERGE_ENABLED` | `true` only after integration rehearsal; otherwise absent or `false` |
| `SDK_PR_APP_BOT_ID` | Numeric GitHub bot user ID of the generator's PR App |
| `SDK_INTEGRATION_APP_BOT_ID` | Numeric GitHub bot user ID of the separate integration App |
| `SDK_INTEGRATION_APP_CLIENT_ID` | Integration App client ID |
| `SDK_REVIEW_RULESET_ID` | Active repository approval ruleset with merger PR-only bypass |
| `SDK_CIRCLECI_STATUS_ACTOR_ID` | Verified numeric identity publishing the required CircleCI status |
| `SDK_RELEASE_ENABLED` | `true` only after the hosted rehearsal succeeds; otherwise absent or `false` |
| `SDK_RELEASE_CONTROL_COMMIT` | Full commit SHA containing the reviewed release workflow and scripts |
| `SDK_RELEASE_SIGNING_PUBLIC_KEY` | Dedicated Ed25519 public signing key |
| `SDK_RELEASE_APP_CLIENT_ID` | Dedicated release App's client ID |

Integration environment secret: `SDK_INTEGRATION_APP_PRIVATE_KEY`.

Release environment secrets: `SDK_RELEASE_APP_PRIVATE_KEY` and
`SDK_RELEASE_SIGNING_PRIVATE_KEY`. Install the release App only on `plaid/plaid-go`
with Contents write, no workflow permission, no Admin and no review bypass.
Choose the approved SSH signing identity separately from App authentication and
verify its signature policy during setup. Do not assume an SSH key can be
registered to an App bot or that a locally valid signature guarantees GitHub's
Verified badge. Generation signatures come from the public spec repository
and are verified using ordinary read-only credentials.

Updating either workflow or any control script requires a reviewed PR and then
updating the control-commit variable. Runs continue to use the pinned scripts; a
workflow that differs from that pin fails closed. The pin does not prevent an
unreviewed workflow from requesting credentials in the first place; required
code-owner review and protected environments remain essential.

## Starting after a manual v49 release

A manual release has no generation manifest. Before the first automated successor,
review and set all three repository variables:

- `SDK_RELEASE_MANUAL_PREDECESSOR_TAG`: exact published manual version, for example `v49.0.0`.
- `SDK_RELEASE_MANUAL_PREDECESSOR_COMMIT`: exact full commit that tag resolves to.
- `SDK_RELEASE_MANUAL_PREDECESSOR_SPEC_COMMIT`: exact public `plaid-openapi` commit used for that release.

Do not infer the spec pin from the SDK version. Record the reviewed public spec
cut as part of the manual release. Missing pins, a different latest tag or commit,
and a candidate spec older than or diverged from that pin all stop publication.
The successor must have full signed generation evidence. After that release,
normal successors use its manifest; the temporary manual pins can be removed.

## Rehearsal and recovery

First bootstrap reviewed v49 content and these release controls with publication
disabled. The generator must use the agreed public spec/config commit, not an older
public spec. Run `workflow_dispatch` on protected `master` with `dry_run=true`.
This explicit manual dry run is permitted while `SDK_RELEASE_ENABLED` remains
false; push-triggered runs and tag publication stay disabled. The workflow derives
`SDK_RELEASE_DRY_RUN` from its event and boolean input; it is not a repository
variable or an alternate publication switch.
By default this builds and packages without creating a real attestation or tag.
Only explicitly setting `attest_dry_run=true` permits a test attestation and tag
preflight; it still cannot publish a tag.

A hosted rehearsal must demonstrate the actual public-generation signature,
cross-repository Apps, protected branch/tag enforcement, release signing and
fresh installation of the canonical archive through a local Go module proxy
(or an approved disposable hosted test tag) before enablement. After the first
real tag, verify a fresh public installation and runtime smoke test before
announcements. Local tests do not establish GitHub OIDC or organization-policy
readiness. Disable both `SDK_AUTOMERGE_ENABLED` and `SDK_RELEASE_ENABLED` to
stop new automated merges and releases. Never repair a published version by moving its tag;
review and publish a new version instead.

Local policy checks use real Git, canonical module ZIPs and SSH signatures, with
GitHub API and signature-verifier boundaries mocked explicitly:

```sh
cd scripts/release/modulezip
GOTOOLCHAIN=go1.26.4 go build -o /tmp/plaid-go-modulezip .
cd ../../..
MODULEZIP_BIN=/tmp/plaid-go-modulezip python3 -m unittest discover -s scripts/release -p 'test_*.py'
actionlint .github/workflows/release.yml .github/workflows/integrate.yml
python3 scripts/release/verify_cli_fixture.py --fixture-dir /tmp/plaid-go-gh-fixture --download
```

`verify_cli_fixture.py` additionally exercises the production verifier with an
official GitHub CLI signed fixture. It checks actual command compatibility and
certificate parsing, not Plaid's still-required hosted rehearsal.
