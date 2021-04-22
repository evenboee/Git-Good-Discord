package abstraction

import (
	"git-good-discord/abstraction/abstraction_interfaces"
	"git-good-discord/discord/discord_interfaces"
	"git-good-discord/gitlab/gitlab_interfaces"
)

var Discord discord_interfaces.Interface
var Gitlab gitlab_interfaces.Interface

type Implementation struct {}

func GetImplementation() abstraction_interfaces.Interface {
	implementation := Implementation{}
	return &implementation
}