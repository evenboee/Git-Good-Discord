package abstraction

import (
	"git-good-discord/abstraction/abstraction_interfaces"
	"git-good-discord/discord/discord_interfaces"
	"git-good-discord/gitlab/gitlab_interfaces"
	"git-good-discord/gitlab/gitlab_structs"
)

var Discord discord_interfaces.Interface
var Gitlab gitlab_interfaces.Interface

type Implementation struct {}

func GetImplementation() abstraction_interfaces.Interface {
	implementation := Implementation{}
	return &implementation
}

func (i Implementation) HandleGitlabMergeRequestNotification(notification gitlab_structs.MergeRequestWebhookNotification, discordChannelID string) {
	// TODO: Handle Merge Request Notification, e.g. send msg in Discord chat
}