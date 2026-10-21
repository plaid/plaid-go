"""Run the exact failure-workflow script against a local GitHub API fake."""
import json
from pathlib import Path
import subprocess
import textwrap
import unittest

ROOT = Path(__file__).resolve().parents[2]
WORKFLOW = ROOT / '.github/workflows/sdk-automation-failures.yml'

class FailureReportingTests(unittest.TestCase):
    def execute(self, *, run_changes=None, owner='ngyna', issues=None):
        workflow = WORKFLOW.read_text()
        script = textwrap.dedent(workflow.split('          script: |\n', 1)[1])
        source = ROOT.name == 'plaid-openapi'
        name = 'Generate Go SDK from approved inputs' if source else 'Release Go SDK'
        path = '.github/workflows/generate-go-sdk.yml' if source else '.github/workflows/release.yml'
        run = dict(name=name, path=path, head_branch='master', head_repository={'full_name':'plaid/'+ROOT.name},
                   event='push', conclusion='failure', id=123, run_attempt=2, display_title='DO-NOT-PUBLISH-RAW-LOGS')
        run.update(run_changes or {})
        payload = dict(script=script, context={'repo':{'owner':'plaid','repo':ROOT.name},'payload':{'workflow_run':run}}, owner=owner, issues=issues or [])
        harness = '''const fs = require('fs');
const p = JSON.parse(fs.readFileSync(0, 'utf8'));
const calls = [];
const github = {rest:{issues:{listForRepo(){},create:async x=>calls.push(['create',x]),update:async x=>calls.push(['update',x])}},
 paginate:async (_,x)=>{calls.push(['list',x]); return p.issues;}};
const AsyncFunction=Object.getPrototypeOf(async function(){}).constructor;
(async()=>{try{await new AsyncFunction('github','context','core','process',p.script)(github,p.context,{info(){}},{env:{SDK_FAILURE_ASSIGNEE:p.owner}});
console.log(JSON.stringify({calls}));}catch(e){console.log(JSON.stringify({calls,error:e.message}));}})();'''
        result = subprocess.run(['node','-e',harness], input=json.dumps(payload),text=True,capture_output=True,check=True)
        return json.loads(result.stdout)

    def test_failure_creates_assigned_issue_without_logs(self):
        result = self.execute()
        self.assertNotIn('error',result)
        fields=result['calls'][-1][1]
        self.assertEqual(result['calls'][-1][0],'create')
        self.assertEqual(fields['assignees'],['ngyna'])
        self.assertIn('/actions/runs/123/attempts/2',fields['body'])
        self.assertNotIn('DO-NOT-PUBLISH',fields['body'])

    def test_updates_one_owned_issue_and_duplicate_delivery_is_noop(self):
        fields=self.execute()['calls'][-1][1]
        issue=dict(number=7,user={'login':'github-actions[bot]'},body=fields['body']+' old',assignees=[{'login':'ngyna'}])
        result=self.execute(issues=[issue])
        self.assertEqual(result['calls'][-1][0],'update')
        self.assertEqual(result['calls'][-1][1]['issue_number'],7)
        issue['body']=fields['body']
        self.assertEqual([x[0] for x in self.execute(issues=[issue])['calls']],['list'])

    def test_untrusted_or_nonfailure_runs_make_no_api_calls(self):
        for change in ({'head_repository':{'full_name':'fork/other'}},{'head_branch':'feature'},
                       {'event':'pull_request'},{'event':'pull_request_target'},{'conclusion':'success'},
                       {'name':'Unrelated workflow'},{'path':'.github/workflows/other.yml'}):
            with self.subTest(change=change): self.assertEqual(self.execute(run_changes=change)['calls'],[])

    def test_sdk_integration_status_trigger_is_reported(self):
        result = self.execute(run_changes={'name':'Integrate generated Go SDK',
                                          'path':'.github/workflows/integrate.yml', 'event':'status'})
        if ROOT.name == 'plaid-go':
            self.assertEqual(result['calls'][-1][0], 'create')
            self.assertIn('"status"', WORKFLOW.read_text().split('    runs-on:', 1)[0])
        else:
            self.assertEqual(result['calls'], [])

    def test_missing_owner_and_invalid_run_identity_fail_before_api(self):
        for args in ({'owner':''},{'owner':'bad user'},{'run_changes':{'id':'unsafe'}},{'run_changes':{'run_attempt':0}}):
            with self.subTest(args=args):
                result=self.execute(**args)
                self.assertIn('error',result)
                self.assertEqual(result['calls'],[])

    def test_user_created_marker_and_pull_request_are_not_modified(self):
        fields=self.execute()['calls'][-1][1]
        for issue in ({'number':7,'user':{'login':'someone'},'body':fields['body']},
                      {'number':8,'user':{'login':'github-actions[bot]'},'body':fields['body'],'pull_request':{}}):
            with self.subTest(issue=issue):self.assertEqual(self.execute(issues=[issue])['calls'][-1][0],'create')

    def test_timeout_is_reported_and_ambiguous_issues_are_not_changed(self):
        self.assertEqual(self.execute(run_changes={'conclusion':'timed_out'})['calls'][-1][0],'create')
        fields=self.execute()['calls'][-1][1]
        issue={'number':7,'user':{'login':'github-actions[bot]'},'body':fields['body']}
        result=self.execute(issues=[issue,dict(issue,number=8)])
        self.assertIn('Multiple SDK failure issues',result['error'])
        self.assertEqual([x[0] for x in result['calls']],['list'])

    def test_reporter_does_not_checkout_or_download_failed_run_code(self):
        workflow=WORKFLOW.read_text()
        self.assertNotIn('actions/checkout',workflow)
        self.assertNotIn('download-artifact',workflow)
        self.assertIn('issues: write',workflow)
        self.assertNotIn('contents: write',workflow)
        self.assertNotIn('secrets.',workflow)

if __name__ == '__main__': unittest.main()
