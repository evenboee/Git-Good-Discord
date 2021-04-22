package gitlab

import (
	"git-good-discord/abstraction/abstraction_interfaces"
	"git-good-discord/gitlab/gitlab_interfaces"
)

var Abstraction abstraction_interfaces.Interface

type Implementation struct {}

func GetImplementation() gitlab_interfaces.Interface {
	implementation := Implementation{}
	return &implementation
}