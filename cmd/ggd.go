package main

import (
	"git-good-discord/abstraction"
	"git-good-discord/discord"
	"git-good-discord/gitlab"
	"git-good-discord/http_serving"
)

func main() {
	// Dependency injection
	abstraction.Discord = discord.GetImplementation()
	abstraction.Gitlab = gitlab.GetImplementation()

	discord.Abstraction = abstraction.GetImplementation()
	gitlab.Abstraction = abstraction.GetImplementation()

	go http_serving.StartWebHandler()
}