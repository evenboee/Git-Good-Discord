package database

import interfaces "git-good-discord/database/database_interfaces"

var Connection interfaces.Database = FirestoreConnection{
	open:   false,
}
