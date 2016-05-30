package app

import (
	"github.com/stretchr/testify/assert"

	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGitHubPing(t *testing.T) {

	server := httptest.NewServer(CreateApp())

	req, _ := http.NewRequest("POST", server.URL+"/webhooks/github/", bytes.NewReader([]byte(PingJson)))
	req.Header.Set("X-GitHub-Event", "ping")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hub-Signature", PingSha)

	if response, err := (&http.Client{}).Do(req); assert.NoError(t, err) {

		if assert.Equal(t, http.StatusOK, response.StatusCode) {

			var resp string

			if err := json.NewDecoder(response.Body).Decode(&resp); assert.NoError(t, err) {

				assert.Equal(t, "pong", resp)
			}
		}
	}
}

const (
	PingSha  = "27577ef3f3d983a7c34c4863cd3a36ad8a3ee128"
	PingJson = `
	{
	  "zen": "Non-blocking is better than blocking.",
	  "hook_id": 8566024,
	  "hook": {
		"type": "Repository",
		"id": 8566024,
		"name": "web",
		"active": true,
		"events": [
		  "push"
		],
		"config": {
		  "content_type": "json",
		  "insecure_ssl": "0",
		  "secret": "********",
		  "url": "http://185.143.172.56/webhooks/github/"
		},
		"updated_at": "2016-05-30T06:52:10Z",
		"created_at": "2016-05-30T06:52:10Z",
		"url": "https://api.github.com/repos/postgres-ci/http200ok/hooks/8566024",
		"test_url": "https://api.github.com/repos/postgres-ci/http200ok/hooks/8566024/test",
		"ping_url": "https://api.github.com/repos/postgres-ci/http200ok/hooks/8566024/pings",
		"last_response": {
		  "code": null,
		  "status": "unused",
		  "message": null
		}
	  },
	  "repository": {
		"id": 56714132,
		"name": "http200ok",
		"full_name": "postgres-ci/http200ok",
		"owner": {
		  "login": "postgres-ci",
		  "id": 16963162,
		  "avatar_url": "https://avatars.githubusercontent.com/u/16963162?v=3",
		  "gravatar_id": "",
		  "url": "https://api.github.com/users/postgres-ci",
		  "html_url": "https://github.com/postgres-ci",
		  "followers_url": "https://api.github.com/users/postgres-ci/followers",
		  "following_url": "https://api.github.com/users/postgres-ci/following{/other_user}",
		  "gists_url": "https://api.github.com/users/postgres-ci/gists{/gist_id}",
		  "starred_url": "https://api.github.com/users/postgres-ci/starred{/owner}{/repo}",
		  "subscriptions_url": "https://api.github.com/users/postgres-ci/subscriptions",
		  "organizations_url": "https://api.github.com/users/postgres-ci/orgs",
		  "repos_url": "https://api.github.com/users/postgres-ci/repos",
		  "events_url": "https://api.github.com/users/postgres-ci/events{/privacy}",
		  "received_events_url": "https://api.github.com/users/postgres-ci/received_events",
		  "type": "Organization",
		  "site_admin": false
		},
		"private": false,
		"html_url": "https://github.com/postgres-ci/http200ok",
		"description": "",
		"fork": false,
		"url": "https://api.github.com/repos/postgres-ci/http200ok",
		"forks_url": "https://api.github.com/repos/postgres-ci/http200ok/forks",
		"keys_url": "https://api.github.com/repos/postgres-ci/http200ok/keys{/key_id}",
		"collaborators_url": "https://api.github.com/repos/postgres-ci/http200ok/collaborators{/collaborator}",
		"teams_url": "https://api.github.com/repos/postgres-ci/http200ok/teams",
		"hooks_url": "https://api.github.com/repos/postgres-ci/http200ok/hooks",
		"issue_events_url": "https://api.github.com/repos/postgres-ci/http200ok/issues/events{/number}",
		"events_url": "https://api.github.com/repos/postgres-ci/http200ok/events",
		"assignees_url": "https://api.github.com/repos/postgres-ci/http200ok/assignees{/user}",
		"branches_url": "https://api.github.com/repos/postgres-ci/http200ok/branches{/branch}",
		"tags_url": "https://api.github.com/repos/postgres-ci/http200ok/tags",
		"blobs_url": "https://api.github.com/repos/postgres-ci/http200ok/git/blobs{/sha}",
		"git_tags_url": "https://api.github.com/repos/postgres-ci/http200ok/git/tags{/sha}",
		"git_refs_url": "https://api.github.com/repos/postgres-ci/http200ok/git/refs{/sha}",
		"trees_url": "https://api.github.com/repos/postgres-ci/http200ok/git/trees{/sha}",
		"statuses_url": "https://api.github.com/repos/postgres-ci/http200ok/statuses/{sha}",
		"languages_url": "https://api.github.com/repos/postgres-ci/http200ok/languages",
		"stargazers_url": "https://api.github.com/repos/postgres-ci/http200ok/stargazers",
		"contributors_url": "https://api.github.com/repos/postgres-ci/http200ok/contributors",
		"subscribers_url": "https://api.github.com/repos/postgres-ci/http200ok/subscribers",
		"subscription_url": "https://api.github.com/repos/postgres-ci/http200ok/subscription",
		"commits_url": "https://api.github.com/repos/postgres-ci/http200ok/commits{/sha}",
		"git_commits_url": "https://api.github.com/repos/postgres-ci/http200ok/git/commits{/sha}",
		"comments_url": "https://api.github.com/repos/postgres-ci/http200ok/comments{/number}",
		"issue_comment_url": "https://api.github.com/repos/postgres-ci/http200ok/issues/comments{/number}",
		"contents_url": "https://api.github.com/repos/postgres-ci/http200ok/contents/{+path}",
		"compare_url": "https://api.github.com/repos/postgres-ci/http200ok/compare/{base}...{head}",
		"merges_url": "https://api.github.com/repos/postgres-ci/http200ok/merges",
		"archive_url": "https://api.github.com/repos/postgres-ci/http200ok/{archive_format}{/ref}",
		"downloads_url": "https://api.github.com/repos/postgres-ci/http200ok/downloads",
		"issues_url": "https://api.github.com/repos/postgres-ci/http200ok/issues{/number}",
		"pulls_url": "https://api.github.com/repos/postgres-ci/http200ok/pulls{/number}",
		"milestones_url": "https://api.github.com/repos/postgres-ci/http200ok/milestones{/number}",
		"notifications_url": "https://api.github.com/repos/postgres-ci/http200ok/notifications{?since,all,participating}",
		"labels_url": "https://api.github.com/repos/postgres-ci/http200ok/labels{/name}",
		"releases_url": "https://api.github.com/repos/postgres-ci/http200ok/releases{/id}",
		"deployments_url": "https://api.github.com/repos/postgres-ci/http200ok/deployments",
		"created_at": "2016-04-20T19:10:09Z",
		"updated_at": "2016-05-17T21:32:50Z",
		"pushed_at": "2016-04-20T20:05:15Z",
		"git_url": "git://github.com/postgres-ci/http200ok.git",
		"ssh_url": "git@github.com:postgres-ci/http200ok.git",
		"clone_url": "https://github.com/postgres-ci/http200ok.git",
		"svn_url": "https://github.com/postgres-ci/http200ok",
		"homepage": null,
		"size": 131,
		"stargazers_count": 0,
		"watchers_count": 0,
		"language": "Go",
		"has_issues": false,
		"has_downloads": true,
		"has_wiki": false,
		"has_pages": false,
		"forks_count": 0,
		"mirror_url": null,
		"open_issues_count": 0,
		"forks": 0,
		"open_issues": 0,
		"watchers": 0,
		"default_branch": "master"
	  },
	  "sender": {
		"login": "kshvakov",
		"id": 978534,
		"avatar_url": "https://avatars.githubusercontent.com/u/978534?v=3",
		"gravatar_id": "",
		"url": "https://api.github.com/users/kshvakov",
		"html_url": "https://github.com/kshvakov",
		"followers_url": "https://api.github.com/users/kshvakov/followers",
		"following_url": "https://api.github.com/users/kshvakov/following{/other_user}",
		"gists_url": "https://api.github.com/users/kshvakov/gists{/gist_id}",
		"starred_url": "https://api.github.com/users/kshvakov/starred{/owner}{/repo}",
		"subscriptions_url": "https://api.github.com/users/kshvakov/subscriptions",
		"organizations_url": "https://api.github.com/users/kshvakov/orgs",
		"repos_url": "https://api.github.com/users/kshvakov/repos",
		"events_url": "https://api.github.com/users/kshvakov/events{/privacy}",
		"received_events_url": "https://api.github.com/users/kshvakov/received_events",
		"type": "User",
		"site_admin": false
	  }
	}
	`
)
