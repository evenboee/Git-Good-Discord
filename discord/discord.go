package discord

import (
	"git-good-discord/abstraction/abstraction_interfaces"
	"git-good-discord/discord/discord_interfaces"
)

var Abstraction abstraction_interfaces.Interface

type Implementation struct {}

func GetImplementation() discord_interfaces.Interface {
	implementation := Implementation{}
	return &implementation
}