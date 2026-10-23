# Automated Go SDK publication

After a reviewed change merges to `plaid/plaid-openapi`, generation produces both
Go clients and a signed record of their exact inputs and content. The integration
workflow checks the generated PR and required CI, then merges it automatically.
The release workflow verifies that merge and generation record, builds and tests
the SDK, attests a canonical Go module ZIP, and publishes a signed version tag.
No human approval or workflow dispatch is required for compatible v49 updates. Both
automatic integration and publication remain disabled until configured and tested.

A separately human-reviewed SDK PR remains an explicit manual fallback. Neither
path may ignore an outstanding changes-requested review or alter an existing tag.

The SDK and `legacy/` must come from the same public OpenAPI spec. The generation
record binds the spec, public config commit, immutable generator images, version,
legacy tree and all tracked SDK content. Public Fern metadata contains only the
CLI, generator and SDK versions; the signed manifest records the shared spec/config commit.

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

Verify that GitHub accepts the CODEOWNERS entries and enable required code-owner
review. Keep required SDK CI and strict up-to-date checks on `master`. Human changes to
controls, config and workflows require review. The dedicated integration App needs
an exception to the PR approval requirement for generated SDK PRs; it must not
bypass required status checks. Its workflow independently verifies the generation
signature, author ID, exact head/base, maintained files and CI identity before
merging. The review exception is attached to the App, not to particular PR
contents; our verifier enforces the generated-content restriction. The current
verifier reads classic branch protection. A rulesets-only setup is not supported
without extending that check, and hosted tests must prove CI cannot be bypassed.
A bot must never post a rubber-stamp approval to satisfy review rules.

Create `sdk-integration` and `release` environments restricted to protected
`master`, with no per-run human approval on the automatic path. The integration
App has Contents write, Pull requests read and Administration read (for checking
branch protection), scoped only to this SDK repository. Its check job obtains a
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
