package main

import (
	"git-good-discord/abstraction"
	"git-good-discord/discord"
	"git-good-discord/gitlab"
	"git-good-discord/http_serving"
	"log"
)

func main() {
	// Select dependency implementations to use
	discordInterface := discord.GetImplementation()
	gitlabInterface := gitlab.GetImplementation()
	abstractionInterface := abstraction.GetImplementation()

	// Dependency injection
	abstraction.Discord = discordInterface
	abstraction.Gitlab = gitlabInterface

	discord.Abstraction = abstractionInterface
	gitlab.Abstraction = abstractionInterface

	http_serving.Gitlab = gitlabInterface

	// Making error channel in case of fatal error
	errorChannel := make(chan error)
	go http_serving.StartWebHandler(errorChannel)
	go discordInterface.Start(errorChannel)

	// Throwing a fatal error and printing it for debugging purposes.
	err := <- errorChannel
	log.Println("A fatal error occured, exiting application:")
	log.Fatal(err)
}