package gitlab_interfaces

import (
	"git-good-discord/gitlab/gitlab_structs"
	"net/http"
)

type Interface interface {
	RegisterWebhook (project gitlab_structs.Project, webhook gitlab_structs.Webhook) (gitlab_structs.WebhookRegistration, error)
    HandleWebhookNotificationHTTP(w http.ResponseWriter, req *http.Request) error
}