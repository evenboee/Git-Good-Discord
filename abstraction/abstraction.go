package abstraction

import (
	"git-good-discord/discord/discord_interfaces"
	"git-good-discord/gitlab/gitlab_interfaces"
	"git-good-discord/gitlab/gitlab_structs"
)

type Implementation struct {
	DiscordService discord_interfaces.Interface
	GitlabService gitlab_interfaces.Interface
}

func (i Implementation) HandleGitlabMergeRequestNotification(notification gitlab_structs.MergeRequestWebhookNotification, discordChannelID string) {
	// TODO: Handle Merge Request Notification, e.g. send msg in Discord chat
}