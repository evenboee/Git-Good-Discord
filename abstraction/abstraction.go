package abstraction

import (
	"git-good-discord/database/database_interfaces"
	"git-good-discord/discord/discord_interfaces"
	"git-good-discord/gitlab/gitlab_interfaces"
)

// Implementation of the abstraction_interfaces.Interface
type Implementation struct {

	// DiscordService required by implementation
	DiscordService discord_interfaces.Interface

	// GitlabService required by implementation
	GitlabService gitlab_interfaces.Interface

	// DatabaseService required by implementation
	DatabaseService database_interfaces.Database
}
