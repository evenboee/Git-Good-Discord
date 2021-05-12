package database_interfaces

import (
	"git-good-discord/database/database_structs"
)

type Database interface {
	ConnectToDatabase(errorChan chan error)
	GetConnection() DatabaseConnection
}

type DatabaseConnection interface {
	AddSubscriber(channel_id string, gitlab_instance string, repo_id string, gitlab_username string, discord_user_id string, issues bool, merge_requests bool) error

	GetSubscribers(channel_id string, gitlab_instance string, repo_id string, gitlab_username string) ([]database_structs.Subscriber, error)

	DeleteSubscriber(channel_id string, gitlab_instance string, repo_id string, gitlab_username string, discord_user_id string) error

	GetChannelSettings(channel_id string) (database_structs.ChannelSettings, error)

	SetChannelPrefix(channel_id string, prefix string) error

	SetChannelLanguage(channel_id string, language string) error

	GetSecurityToken(channel_id string, gitlab_instance string, repo_id string) (string, error)

	AddSecurityToken(channel_id string, gitlab_instance string, repo_id string, token string) error

	AddAccessToken(channel_id string, gitlab_instance string, repo_id string, token string) error

	GetAccessToken(channel_id string, gitlab_instance string, repo_id string) (string, error)

	Close() error
}
