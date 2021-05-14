package gitlab

import (
	"git-good-discord/abstraction/abstraction_interfaces"
	"git-good-discord/database/database_interfaces"
)

// Implementation is the implementation for gitlab_interfaces.Interface
type Implementation struct {

	// AbstractionService required by implementation
	AbstractionService abstraction_interfaces.Interface

	// DatabaseService required by implementation
	DatabaseService database_interfaces.Database

}
