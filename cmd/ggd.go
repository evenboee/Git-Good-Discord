package main

import (
	"git-good-discord/abstraction"
	"git-good-discord/discord"
	"git-good-discord/gitlab"
	"git-good-discord/http_serving"
	"log"
)

func main() {
	// Dependency injection
	abstraction.Discord = discord.GetImplementation()
	abstraction.Gitlab = gitlab.GetImplementation()

	discord.Abstraction = abstraction.GetImplementation()
	gitlab.Abstraction = abstraction.GetImplementation()

	// Making error channel in case of fatal error
	errorChannel := make(chan error)
	go http_serving.StartWebHandler(errorChannel)

	// Throwing a fatal error and printing it for debugging purposes.
	err := <- errorChannel
	log.Println("A fatal error occured, exiting application:")
	log.Fatal(err)
}