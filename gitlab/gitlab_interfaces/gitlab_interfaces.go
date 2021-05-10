package gitlab_interfaces

import (
	"git-good-discord/gitlab/gitlab_structs"
	"net/http"
)

type Interface interface {
	RegisterWebhook (project gitlab_structs.Project, webhook gitlab_structs.Webhook) (gitlab_structs.WebhookRegistration, error)
    HandleWebhookNotificationHTTP(w http.ResponseWriter, req *http.Request) error

	// GetUserByUsername gets user information using username.
	// only scheme and host will be extracted from instanceURL
	GetUserByUsername(instanceURL string, username string) (gitlab_structs.User, error)

	// GetUserByID gets user information using id.
	// only scheme and host will be extracted from instanceURL
	GetUserByID(instanceURL string, id int) (gitlab_structs.User, error)

}