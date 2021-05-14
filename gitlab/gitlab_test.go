package gitlab

import (
	"bytes"
	"git-good-discord/abstraction/abstraction_mocks"
	"git-good-discord/database/database_mocks"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/url"
	"testing"
)

type ClosingBuffer struct {
	*bytes.Buffer
}

func (cb *ClosingBuffer) Close() (err error) {
	return
}

func TestImplementation_HandleWebhookNotificationHTTP(t *testing.T) {
	mockctrl := gomock.NewController(t)
	defer mockctrl.Finish()

	abstractionMock := abstraction_mocks.NewMockInterface(mockctrl)
	databaseMock := database_mocks.NewMockDatabase(mockctrl)
	databaseConnectionMock := database_mocks.NewMockDatabaseConnection(mockctrl)

	implementation := Implementation{
		AbstractionService: abstractionMock,
		DatabaseService:    databaseMock,
	}

	// Test JSON emulating notification HTTP body
	testJSON := `{ "object_kind": "issue", "event_type": "issue", "user": { "id": 440, "name": "Ruben Christoffer Hegland-Antonsen", "username": "ruben", "avatar_url": "https://secure.gravatar.com/avatar/d7df274e822e55e8e09e5b6874b11007?s=80&d=identicon", "email": "rubench@stud.ntnu.no" }, "project": { "id": 1965, "name": "testing - Git GOod Discord - Group 1", "description": "", "web_url": "https://git.gvk.idi.ntnu.no/simen_bai/testing-git-good-discord-group-1", "avatar_url": null, "git_ssh_url": "git@git.gvk.idi.ntnu.no:simen_bai/testing-git-good-discord-group-1.git", "git_http_url": "https://git.gvk.idi.ntnu.no/simen_bai/testing-git-good-discord-group-1.git", "namespace": "Simen Bai", "visibility_level": 20, "path_with_namespace": "simen_bai/testing-git-good-discord-group-1", "default_branch": "master", "ci_config_path": null, "homepage": "https://git.gvk.idi.ntnu.no/simen_bai/testing-git-good-discord-group-1", "url": "git@git.gvk.idi.ntnu.no:simen_bai/testing-git-good-discord-group-1.git", "ssh_url": "git@git.gvk.idi.ntnu.no:simen_bai/testing-git-good-discord-group-1.git", "http_url": "https://git.gvk.idi.ntnu.no/simen_bai/testing-git-good-discord-group-1.git" }, "object_attributes": { "author_id": 440, "closed_at": null, "confidential": false, "created_at": "2021-05-12 18:49:33 UTC", "description": "mmm", "discussion_locked": null, "due_date": null, "id": 7327, "iid": 6, "last_edited_at": null, "last_edited_by_id": null, "milestone_id": null, "moved_to_id": null, "duplicated_to_id": null, "project_id": 1965, "relative_position": null, "state_id": 1, "time_estimate": 0, "title": "mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm", "updated_at": "2021-05-12 18:49:33 UTC", "updated_by_id": null, "weight": null, "url": "https://git.gvk.idi.ntnu.no/simen_bai/testing-git-good-discord-group-1/-/issues/6", "total_time_spent": 0, "human_total_time_spent": null, "human_time_estimate": null, "assignee_ids": [ 440 ], "assignee_id": 440, "labels": [], "state": "opened", "action": "open" }, "labels": [], "changes": { "author_id": { "previous": null, "current": 440 }, "created_at": { "previous": null, "current": "2021-05-12 18:49:33 UTC" }, "description": { "previous": null, "current": "mmm" }, "id": { "previous": null, "current": 7327 }, "iid": { "previous": null, "current": 6 }, "project_id": { "previous": null, "current": 1965 }, "title": { "previous": null, "current": "mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm" }, "updated_at": { "previous": null, "current": "2021-05-12 18:49:33 UTC" } }, "repository": { "name": "testing - Git GOod Discord - Group 1", "url": "git@git.gvk.idi.ntnu.no:simen_bai/testing-git-good-discord-group-1.git", "description": "", "homepage": "https://git.gvk.idi.ntnu.no/simen_bai/testing-git-good-discord-group-1" }, "assignees": [ { "id": 440, "name": "Ruben Christoffer Hegland-Antonsen", "username": "ruben", "avatar_url": "https://secure.gravatar.com/avatar/d7df274e822e55e8e09e5b6874b11007?s=80&d=identicon", "email": "rubench@stud.ntnu.no" } ] }`

	req := http.Request{}
	req.Body = &ClosingBuffer{bytes.NewBufferString(testJSON)}
	req.URL = &url.URL{
		Scheme: "https",
		Host:   "git.gvk.idi.ntnu.no",
		Path:   "/gitlab/831517371185102848",
	}

	req.Header = http.Header{}

	// Expected invocations
	databaseMock.EXPECT().GetConnection().Return(databaseConnectionMock)
	databaseConnectionMock.EXPECT().GetSecurityToken("831517371185102848", "git.gvk.idi.ntnu.no", "1965").Times(1)

	abstractionMock.EXPECT().HandleGitlabNotification(gomock.Any(), "831517371185102848").Times(1)

	implementation.HandleWebhookNotificationHTTP(nil, &req)
}
