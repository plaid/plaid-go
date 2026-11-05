#!/usr/bin/env python3
"""Write the required status from trusted controls; never execute PR code."""
import argparse
import os
from pathlib import Path
import sys
import tempfile

import integrate
import policy
from policy import SDK_REPO, ReleaseError, api, canonical, decimal, git, hex_value, parse_json, require, run


class StatusWriteError(ReleaseError):
    pass


def report(head, state):
    hex_value(head, 40, "verification status head")
    require(state in ("pending", "failure", "success"), "Unexpected verification status")
    token = os.environ.get("SDK_VERIFICATION_STATUS_TOKEN", "")
    require(token, "Missing trusted workflow status token")
    run_id = os.environ.get("GITHUB_RUN_ID", "")
    decimal(run_id, "verification workflow run ID")
    # Fixed descriptions: never copy PR content, tool errors or credentials into statuses.
    body = canonical({"state": state, "context": integrate.VERIFIED_CONTEXT,
                      "description": "Trusted SDK generation or exact-head human review verification",
                      "target_url": f"https://github.com/{SDK_REPO}/actions/runs/{run_id}"})
    try:
        run("gh", "api", "--hostname", "github.com", "--method", "POST",
            f"repos/{SDK_REPO}/statuses/{head}", "--input", "-", input=body,
            env=dict(os.environ, GH_TOKEN=token, GH_HOST="github.com"))
    except ReleaseError as error:
        raise StatusWriteError("Could not update required verification status") from error


def human_approval(pull):
    reviews = policy.human_reviews(pull)
    if not any(item.get("state") == "APPROVED" and item.get("commit_id") == pull["head"]["sha"]
               and type(item.get("user", {}).get("id")) is int
               and item["user"]["id"] != pull.get("user", {}).get("id") for item in reviews):
        raise integrate.NotReady("Waiting for human approval of the exact PR head")
    query = ('query($number:Int!){repository(owner:"plaid",name:"plaid-go"){'
             'pullRequest(number:$number){headRefOid baseRefName reviewDecision state isDraft}}}')
    data = parse_json(run("gh", "api", "--hostname", "github.com", "graphql", "-f", "query=" + query,
                          "-F", "number=" + str(pull["number"])))
    require(not data.get("errors"), "Cannot establish GitHub review decision")
    current = data.get("data", {}).get("repository", {}).get("pullRequest") or {}
    require(current.get("headRefOid") == pull["head"]["sha"] and current.get("baseRefName") == "master"
            and current.get("state") == "OPEN" and current.get("isDraft") is False,
            "Human-reviewed PR changed during verification")
    if current.get("reviewDecision") != "APPROVED":
        raise integrate.NotReady("Waiting for GitHub's required review decision")


def open_pull(pull):
    require(pull.get("state") == "open" and pull.get("draft") is False and not pull.get("merged"),
            "Verification requires an open ready PR")
    require(pull.get("base", {}).get("repo", {}).get("full_name") == SDK_REPO
            and pull["base"].get("ref") == "master", "Unexpected PR target")
    for side in ("base", "head"):
        hex_value(pull.get(side, {}).get("sha"), 40, f"PR {side}")


def verify_one(control, control_commit, number):
    pull = api(f"repos/{SDK_REPO}/pulls/{number}")
    head = pull.get("head", {}).get("sha")
    hex_value(head, 40, "PR status head")
    # Clear a previous success before any signature, CI or review check can fail.
    report(head, "pending")
    try:
        open_pull(pull)
        require(integrate.current_base() == pull["base"]["sha"], "PR base moved")
        integrate.strict_status_protection()
        integrate.approval_ruleset()
        try:
            policy.require_generated_pull(pull)
            generated = True
        except ReleaseError:
            generated = False
        with tempfile.TemporaryDirectory(prefix="sdk-pr-verification-") as temporary:
            repo = Path(temporary) / "objects.git"
            run("git", "init", "--bare", repo)
            if generated:
                integrate.check(repo, control, control_commit, number,
                                expected_base=pull["base"]["sha"], expected_head=head, status_only=True)
            else:
                # Read only trusted base/workflow Git objects, never fork or PR code.
                integrate.fetch_objects(repo, pull["base"]["sha"], os.environ["GITHUB_SHA"])
                expected = git(control, "show", f"{control_commit}:{integrate.WORKFLOW}")
                require(git(repo, "show", f"{os.environ['GITHUB_SHA']}:{integrate.WORKFLOW}") == expected
                        and git(repo, "show", f"{pull['base']['sha']}:{integrate.WORKFLOW}") == expected,
                        "Verification workflow differs from pinned controls")
                integrate.ci_and_reviews(number, head)
                human_approval(pull)
        latest = api(f"repos/{SDK_REPO}/pulls/{number}")
        open_pull(latest)
        require(latest["head"]["sha"] == head and latest["base"]["sha"] == pull["base"]["sha"],
                "PR changed before verification status publication")
    except integrate.NotReady:
        return  # Pending remains non-mergeable, including after a prior success.
    except Exception:
        report(head, "failure")
        raise
    report(head, "success")


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--control", type=Path, required=True)
    parser.add_argument("--control-commit", required=True)
    args = parser.parse_args()
    integrate.context(args.control, args.control_commit, status_only=True)
    if os.environ.get("SDK_INTEGRATION_DRY_RUN") == "true":
        print("Dry run: no verification statuses written")
        return
    pulls = api(f"repos/{SDK_REPO}/pulls?state=open&base=master&per_page=100", paginate=True)
    errors = 0
    for pull in pulls:
        try:
            verify_one(args.control, args.control_commit, pull["number"])
        except StatusWriteError:
            raise
        except (ReleaseError, OSError, ValueError, KeyError):
            errors += 1
    # An unrelated manual PR awaiting fixes must not block a valid generated PR.
    # Its own required status remains failed; generated integration rechecks itself.
    print(f"Verification completed; {errors} PRs have failing verification statuses")


if __name__ == "__main__":
    try:
        main()
    except (ReleaseError, OSError, ValueError, KeyError) as error:
        sys.exit(str(error))
