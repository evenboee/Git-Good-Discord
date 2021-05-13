package database

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"errors"
	firebase "firebase.google.com/go"
	"git-good-discord/database/database_interfaces"
	"git-good-discord/database/database_structs"
	"git-good-discord/utils"
	"google.golang.org/api/iterator"
)

type FirestoreDatabase struct{}

type FirestoreConnection struct {
	open   bool
	ctx    context.Context
	client *firestore.Client
}

const Subscribers = "subscribers"
const Instance = "instances"
const Repos = "repos"
const Channels = "channels"

var connectionNotOpenError = errors.New("firestore connection is not open")

func (db FirestoreDatabase) ConnectToDatabase(errorChan chan error) {
	ctx := context.Background()
	credentials := utils.GetFirestore()

	app, err := firebase.NewApp(ctx, nil, credentials)
	if err != nil {
		errorChan <- err
		return
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		errorChan <- err
		return
	}

	databaseConnection = FirestoreConnection{
		open:   true,
		ctx:    ctx,
		client: client,
	}
}

func (db FirestoreDatabase) GetConnection() database_interfaces.DatabaseConnection {
	return databaseConnection
}
func (conn FirestoreConnection) AddSubscriber(channel_id string, gitlab_instance string, repo_id string, gitlab_username string, discord_user_id string, issues bool, merge_requests bool) error {
	if conn.open != true {
		return connectionNotOpenError
	}
	subscriber := database_structs.Subscriber{
		Issues:        issues,
		MergeRequests: merge_requests,
	}
	// Navigating to resource: channels/{channel_id}/instances/{gitlab_instance}/repos/{repo_id}/subscribers/{gitlab_username}/subscribers/{discord_user_id}/
	collection := conn.client.Collection(Channels).Doc(channel_id).Collection(Instance).Doc(gitlab_instance).Collection(Repos).Doc(repo_id).Collection(Subscribers).Doc(gitlab_username).Collection(Subscribers)
	_, err := collection.Doc(discord_user_id).Set(conn.ctx, subscriber)
	if err != nil {
		return err
	}

	return nil
}

func (conn FirestoreConnection) GetSubscribers(channel_id string, gitlab_instance string, repo_id string, gitlab_username string) ([]database_structs.Subscriber, error) {
	if conn.open != true {
		return nil, connectionNotOpenError
	}
	// Navigating to resource: channels/{channel_id}/instances/{gitlab_instance}/repos/{repo_id}/subscribers/{gitlab_username}/subscribers/{discord_user_id}/
	iter := conn.client.Collection(Channels).Doc(channel_id).Collection(Instance).Doc(gitlab_instance).Collection(Repos).Doc(repo_id).Collection(Subscribers).Doc(gitlab_username).Collection(Subscribers).Documents(conn.ctx)
	var subscribers []database_structs.Subscriber
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var subscriber database_structs.Subscriber

		k, err := json.Marshal(doc.Data())
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(k, &subscriber)
		if err != nil {
			return nil, err
		}

		subscriber.DiscordUserId = doc.Ref.ID
		subscribers = append(subscribers, subscriber)
	}

	return subscribers, nil
}

func (conn FirestoreConnection) DeleteSubscriber(channel_id, gitlab_instance string, repo_id string, gitlab_username string, discord_user_id string) error {
	if conn.open != true {
		return connectionNotOpenError
	}
	// Navigating to resource: channels/{channel_id}/instances/{gitlab_instance}/repos/{repo_id}/subscribers/{gitlab_username}/subscribers/{discord_user_id}/
	subscriber := conn.client.Collection(Channels).Doc(channel_id).Collection(Instance).Doc(gitlab_instance).Collection(Repos).Doc(repo_id).Collection(Subscribers).Doc(gitlab_username).Collection(Subscribers).Doc(discord_user_id)
	_, err := subscriber.Delete(conn.ctx)
	return err
}

func (conn FirestoreConnection) Close() error {
	if conn.open != true {
		return connectionNotOpenError
	}
	return conn.client.Close()
}
