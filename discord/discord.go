package discord

import (
	"git-good-discord/abstraction/abstraction_interfaces"
	"git-good-discord/database/database_interfaces"
	"git-good-discord/gitlab/gitlab_interfaces"
)

type Implementation struct {
	AbstractionService abstraction_interfaces.Interface
	DatabaseService    database_interfaces.Database
	GitlabService    gitlab_interfaces.Interface
}