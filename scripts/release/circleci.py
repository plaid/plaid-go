"""Verify the latest upstream CircleCI job for a Go SDK candidate."""

from contextlib import closing
from http.client import HTTPSConnection, HTTPException
from urllib.parse import quote

from policy import ReleaseError, hex_value, parse_json, require


def verify_build(branch, head):
    """Require CircleCI's latest job on this branch to have built the exact head."""
    require(isinstance(branch, str) and branch not in ("", ".", ".."), "Missing CircleCI branch")
    hex_value(head, 40, "CircleCI candidate commit")
    # The branch jobs endpoint orders jobs by build number. Leave it unfiltered:
    # a newer pending or failed job must supersede an earlier success.
    path = f"/api/v1.1/project/github/plaid/plaid-go/tree/{quote(branch, safe='')}?limit=1"
    limit = 1024 * 1024
    try:
        with closing(HTTPSConnection("circleci.com", timeout=10)) as connection:
            connection.request("GET", path, headers={"Accept": "application/json", "Cache-Control": "no-cache"})
            response = connection.getresponse()
            require(response.status == 200, f"CircleCI returned HTTP {response.status}")
            payload = response.read(limit + 1)
            require(len(payload) <= limit, "CircleCI job response is too large")
            jobs = parse_json(payload)
    except (HTTPException, OSError, ValueError, RecursionError) as exc:
        raise ReleaseError("Could not read the latest CircleCI job") from exc

    require(isinstance(jobs, list) and len(jobs) == 1 and isinstance(jobs[0], dict),
            "CircleCI has no confirmed latest job for this branch")
    job = jobs[0]
    require(job.get("vcs_url") == "https://github.com/plaid/plaid-go" and
            job.get("branch") == branch and job.get("vcs_revision") == head,
            "Latest CircleCI job does not match the repository, branch and commit")
    workflow = job.get("workflows")
    require(isinstance(workflow, dict) and workflow.get("workflow_name") == "build" and
            workflow.get("job_name") == "build", "Latest CircleCI job is not the required build")
    require(job.get("why") == "github" and job.get("vcs_tag") is None and
            job.get("build_parameters") == {} and job.get("ssh_users") == [],
            "CircleCI build was not a normal VCS-triggered branch job")
    require(job.get("lifecycle") == "finished" and job.get("status") == "success" and
            job.get("outcome") == "success", "Latest CircleCI build has not succeeded")
