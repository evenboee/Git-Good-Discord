package abstraction

import (
	"git-good-discord/database/database_interfaces"
	"git-good-discord/discord/discord_interfaces"
	"git-good-discord/gitlab/gitlab_interfaces"
)

type Implementation struct {
	DiscordService discord_interfaces.Interface
	GitlabService gitlab_interfaces.Interface
	DatabaseService database_interfaces.Database
}