package gitlab_interfaces

import (
	"git-good-discord/gitlab/gitlab_structs"
	"net/http"
)

// Interface for Gitlab
type Interface interface {

	// RegisterWebhook will register a Webhook for the given Gitlab Project and
	// return Registration information. This function does not store the registration
	// anywhere, nor does it create some sort of webhook invocation handler. That
	// kind of functionality is up to the caller to implement.
	RegisterWebhook(project gitlab_structs.Project, webhook gitlab_structs.Webhook) (gitlab_structs.WebhookRegistration, error)

	// GetRegisteredWebhooks gets all the registered webhooks for the given project
	GetRegisteredWebhooks(project gitlab_structs.Project) ([]gitlab_structs.WebhookRegistration, error)

	// DoesWebhookWithURLExist checks if a webhook with the given invocation URL has been registered
	DoesWebhookWithURLExist(project gitlab_structs.Project, invocationURL string) (bool, error)

	// GetWebhookInvocationURL gets the properly formatted webhook invocation URL given discord channel ID
	GetWebhookInvocationURL(discordChannelID string) (string, error)

	// HandleWebhookNotificationHTTP handles incoming gitlab notifications
	HandleWebhookNotificationHTTP(w http.ResponseWriter, req *http.Request) error

	// GetUserByUsername gets user information using username.
	// only scheme and host will be extracted from instanceURL
	GetUserByUsername(instanceURL string, username string) (gitlab_structs.User, error)

	// GetUserByID gets user information using id.
	// only scheme and host will be extracted from instanceURL
	GetUserByID(instanceURL string, id int) (gitlab_structs.User, error)
}
