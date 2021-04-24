package discord_interfaces

import "git-good-discord/discord/discord_structs"

type Interface interface {
	Start(chan error)
	SendMessage(discord_structs.Message) error
}