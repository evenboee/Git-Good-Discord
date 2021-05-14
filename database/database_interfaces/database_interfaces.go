package database_interfaces

import (
	"git-good-discord/database/database_structs"
)

// Database contains functions for interacting with a database
type Database interface {

	// ConnectToDatabase will connect to the database and write an error to error
	// channel if something goes wrong
	ConnectToDatabase(errorChan chan error)

	// GetConnection will get the current database connection for this database
	GetConnection() DatabaseConnection

}

// DatabaseConnection contains functions for interacting with an open database
// connection
type DatabaseConnection interface {

	// AddSubscriber will add a subscriber for a given gitlab username
	AddSubscriber(channel_id string, gitlab_instance string, repo_id string, gitlab_username string, discord_user_id string, issues bool, merge_requests bool) error

	// GetSubscribers will get subscribers for a given gitlab username
	GetSubscribers(channel_id string, gitlab_instance string, repo_id string, gitlab_username string) ([]database_structs.Subscriber, error)

	// DeleteSubscriber will delete a subscriber for a given gitlab username
	DeleteSubscriber(channel_id string, gitlab_instance string, repo_id string, gitlab_username string, discord_user_id string) error

	// GetChannelSettings will get channel settings for a given channel
	GetChannelSettings(channel_id string) (database_structs.ChannelSettings, error)

	// SetChannelPrefix will set the command prefix for a given channel
	SetChannelPrefix(channel_id string, prefix string) error

	// SetChannelLanguage will set the language for the given channel
	SetChannelLanguage(channel_id string, language string) error

	// GetSecurityToken will get the security token for the given gitlab project
	GetSecurityToken(channel_id string, gitlab_instance string, repo_id string) (string, error)

	// AddSecurityToken will add a security token for the given gitlab project
	AddSecurityToken(channel_id string, gitlab_instance string, repo_id string, token string) error

	// AddAccessToken will add an access token for the given gitlab project
	AddAccessToken(channel_id string, gitlab_instance string, repo_id string, token string) error

	// GetAccessToken will get the access token for the given gitlab project
	GetAccessToken(channel_id string, gitlab_instance string, repo_id string) (string, error)

	// GetAllSubscriptions gets all subscriptions of a user for a given channel
	GetAllSubscriptions(channel_id string, discord_user_id string) ([]database_structs.Subscriber, error)

	// Close will close the database connection
	Close() error
}
