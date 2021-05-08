package database_interfaces

import "git-good-discord/database/database_structs"

type Database interface {
	AddSubscriber(channel_id string, gitlab_instance string, repo_id string, gitlab_username string, discord_user_id string, issues bool, merge_requests bool) error

	GetSubscribers(channel_id string, gitlab_instance string, repo_id string, gitlab_username string) ([]database_structs.Subscriber, error)

	DeleteSubscriber(channel_id string, gitlab_instance string, repo_id string, gitlab_username string, discord_user_id string) error

	Close() error
}
