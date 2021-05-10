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

func InjectAndInitializeServices () {
	// Inject
	gitlabService = &gitlabImplementation
	discordService = &discordImplementation
	abstractionService = &abstractionImplementation
	databaseService = &databaseImplementation
	webHandlerService = &webHandlerImplementation

	// Initialize dependencies
	gitlabImplementation.AbstractionService = abstractionService

	discordImplementation.AbstractionService = abstractionService
	discordImplementation.DatabaseService = databaseService

	abstractionImplementation.DiscordService = discordService
	abstractionImplementation.GitlabService = gitlabService

	webHandlerImplementation.GitlabService = gitlabService
}

func GetGitlab() gitlab_interfaces.Interface {
	return gitlabService
}

func GetDiscord() discord_interfaces.Interface {
	return discordService
}

func GetAbstraction() abstraction_interfaces.Interface {
	return abstractionService
}

func GetDatabase () database_interfaces.Database{
	return databaseService
}

func GetWebHandler () http_serving_interfaces.WebHandler {
	return webHandlerService
}