package main

import (
	"git-good-discord/discord/discord_messages"
	"git-good-discord/services"
	"log"
)

func main() {
	// Dependency injection
	services.InjectAndInitializeServices()

	// Making error channel in case of fatal error
	errorChannel := make(chan error)

	// Connect to database
	services.GetDatabase().ConnectToDatabase(errorChannel)
	defer services.GetDatabase().GetConnection().Close()

	// Load language packs
	discord_messages.LoadLanguageFiles(errorChannel)

	// Start go-routines
	go services.GetWebHandler().Start(errorChannel)
	go services.GetDiscord().Start(errorChannel)

	// Throwing a fatal error and printing it for debugging purposes.
	err := <- errorChannel
	log.Println("A fatal error occured, exiting application:")
	log.Fatalln(err)
}