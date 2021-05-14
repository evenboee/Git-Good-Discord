package abstraction_interfaces

import "git-good-discord/gitlab/gitlab_structs"

// Interface for abstraction layer
type Interface interface {

	// HandleGitlabNotification handles a gitlab notification
	// once it has been unmarshalled and verified
	HandleGitlabNotification(notification gitlab_structs.WebhookNotification, discordChannelID string)

}