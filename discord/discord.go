package discord

import (
	"git-good-discord/abstraction/abstraction_interfaces"
	"git-good-discord/database/database_interfaces"
	"git-good-discord/gitlab/gitlab_interfaces"
)

// Implementation of the discord_interfaces.Interface
type Implementation struct {

	// AbstractionService required by implementation
	AbstractionService abstraction_interfaces.Interface

	// DatabaseService required by implementation
	DatabaseService database_interfaces.Database

	// GitlabService required by implementation
	GitlabService gitlab_interfaces.Interface
}
