package discord_interfaces

import "git-good-discord/discord/discord_structs"

type Interface interface {
	Start(chan error)
	SendMessage(message discord_structs.EmbeddedMessage) error

}