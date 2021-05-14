package discord_interfaces

import "git-good-discord/discord/discord_structs"

// Interface for the discord bot
type Interface interface {

	// Start starts the discord bot
	Start(chan error)

	// SendMessage sends a message on discord
	SendMessage(message discord_structs.EmbeddedMessage) error

}