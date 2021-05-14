package database

import (
	"cloud.google.com/go/firestore"
	"encoding/json"
	"git-good-discord/database/database_structs"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (conn FirestoreConnection) SetChannelPrefix(channel_id string, prefix string) error {
	return setChannelField(conn, channel_id, "prefix", prefix)
}

func (conn FirestoreConnection) SetChannelLanguage(channel_id string, language string) error {
	return setChannelField(conn, channel_id, "language", language)
}

func (conn FirestoreConnection) GetChannelSettings(channel_id string) (database_structs.ChannelSettings, error) {
	channelSettings := database_structs.ChannelSettings{}
	if conn.open != true {
		return channelSettings, errConnectionNotOpen
	}
	doc, err := conn.client.Collection(Channels).Doc(channel_id).Get(conn.ctx)

	if err != nil {
		if status.Code(err) == codes.NotFound {
			err = nil
		}
		return channelSettings, err
	}

	val, err := json.Marshal(doc.Data())
	if err != nil {
		return channelSettings, err
	}

	err = json.Unmarshal(val, &channelSettings)
	return channelSettings, err
}

// setChannelField will set a firestore field with the given value
func setChannelField(conn FirestoreConnection, channel_id string, fieldName string, fieldValue interface{}) error {
	if conn.open != true {
		return errConnectionNotOpen
	}
	_, err := conn.client.Collection(Channels).Doc(channel_id).Set(conn.ctx, map[string]interface{}{
		fieldName: fieldValue,
	}, firestore.MergeAll) // Document might not exist. MergeAll ensures changes are merged to avoid overwriting entire documents
	// https://firebase.google.com/docs/firestore/manage-data/add-data#set_a_document
	return err
}
