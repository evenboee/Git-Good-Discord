package abstraction_interfaces

import "git-good-discord/gitlab/gitlab_structs"

type Interface interface {
	HandleGitlabNotification(notification gitlab_structs.WebhookNotification, discordChannelID string)
}