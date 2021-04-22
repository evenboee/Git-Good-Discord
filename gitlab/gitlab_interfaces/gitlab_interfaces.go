package gitlab_interfaces

import "git-good-discord/gitlab/gitlab_structs"

type Interface interface {
	RegisterWebhook (project gitlab_structs.Project, webhook gitlab_structs.Webhook) (gitlab_structs.WebhookRegistration, error)
}