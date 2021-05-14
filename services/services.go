package services

import (
	"git-good-discord/abstraction"
	"git-good-discord/abstraction/abstraction_interfaces"
	"git-good-discord/database"
	"git-good-discord/database/database_interfaces"
	"git-good-discord/discord"
	"git-good-discord/discord/discord_interfaces"
	"git-good-discord/gitlab"
	"git-good-discord/gitlab/gitlab_interfaces"
	"git-good-discord/http_serving"
	"git-good-discord/http_serving/http_serving_interfaces"
)

// Services
var (
	gitlabService gitlab_interfaces.Interface
	discordService discord_interfaces.Interface
	abstractionService abstraction_interfaces.Interface
	databaseService database_interfaces.Database
	webHandlerService http_serving_interfaces.WebHandler
)

// Implementations
var (
	gitlabImplementation gitlab.Implementation
	discordImplementation discord.Implementation
	abstractionImplementation abstraction.Implementation
	databaseImplementation database.FirestoreDatabase
	webHandlerImplementation http_serving.Implementation
)

// InjectAndInitializeServices will perform dependency injection for all the
// services and initialize their dependencies so that they can be used
func InjectAndInitializeServices () {
	// Inject
	gitlabService = &gitlabImplementation
	discordService = &discordImplementation
	abstractionService = &abstractionImplementation
	databaseService = &databaseImplementation
	webHandlerService = &webHandlerImplementation

	// Initialize dependencies
	gitlabImplementation.AbstractionService = abstractionService
	gitlabImplementation.DatabaseService = databaseService

	discordImplementation.AbstractionService = abstractionService
	discordImplementation.DatabaseService = databaseService
	discordImplementation.GitlabService = gitlabService

	abstractionImplementation.DiscordService = discordService
	abstractionImplementation.GitlabService = gitlabService
	abstractionImplementation.DatabaseService = databaseService

	webHandlerImplementation.GitlabService = gitlabService
}

// GetGitlab gets the Gitlab service explicitly without being injected
func GetGitlab() gitlab_interfaces.Interface {
	return gitlabService
}

// GetDiscord gets the Discord service explicitly without being injected
func GetDiscord() discord_interfaces.Interface {
	return discordService
}

// GetAbstraction gets the Abstraction service explicitly without being injected
func GetAbstraction() abstraction_interfaces.Interface {
	return abstractionService
}

// GetDatabase gets the Database service explicitly without being injected
func GetDatabase () database_interfaces.Database{
	return databaseService
}

// GetWebHandler gets the Web Handler service explicitly without being injected
func GetWebHandler () http_serving_interfaces.WebHandler {
	return webHandlerService
}