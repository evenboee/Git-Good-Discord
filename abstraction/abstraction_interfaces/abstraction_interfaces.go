package abstraction_interfaces

import "git-good-discord/gitlab/gitlab_structs"

type Interface interface {
	HandleGitlabMergeRequestNotification(notification gitlab_structs.MergeRequestWebhookNotification)
}