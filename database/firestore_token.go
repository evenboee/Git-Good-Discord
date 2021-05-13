package database

import (
	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func addTokenToFirestore(conn FirestoreConnection, channel_id string, gitlab_instance string, repo_id string, token string, token_type string) error {
	if conn.open != true {
		return connectionNotOpenError
	}

	_, err := conn.client.Collection(Channels).Doc(channel_id).Collection(Instance).Doc(gitlab_instance).Collection(Repos).Doc(repo_id).Set(conn.ctx, map[string]interface{}{
		token_type: token,
	}, firestore.MergeAll) // Document might not exist. MergeAll ensures changes are merged to avoid overwriting entire documents

	return err
}

func getTokenFromFirestore(conn FirestoreConnection, channel_id string, gitlab_instance string, repo_id string, token_name string) (string, error) {
	if conn.open != true {
		return "", connectionNotOpenError
	}

	dsnap, err := conn.client.Collection(Channels).Doc(channel_id).Collection(Instance).Doc(gitlab_instance).Collection(Repos).Doc(repo_id).Get(conn.ctx)
	if err != nil {
		return "", err
	}

	data := dsnap.Data()
	if v, ok := data[token_name]; ok {
		return v.(string), nil
	} else {
		return "", nil
	}
}

func (conn FirestoreConnection) AddSecurityToken(channel_id string, gitlab_instance string, repo_id string, token string) error {
	return addTokenToFirestore(conn, channel_id, gitlab_instance, repo_id, token, "security_token")
}

func (conn FirestoreConnection) AddAccessToken(channel_id string, gitlab_instance string, repo_id string, token string) error {
	return addTokenToFirestore(conn, channel_id, gitlab_instance, repo_id, token, "access_token")
}

func (conn FirestoreConnection) GetAccessToken(channel_id string, gitlab_instance string, repo_id string) (string, error) {
	return getTokenFromFirestore(conn, channel_id, gitlab_instance, repo_id, "access_token")
}

func (conn FirestoreConnection) GetSecurityToken(channel_id string, gitlab_instance string, repo_id string) (string, error) {
	token, err := getTokenFromFirestore(conn, channel_id, gitlab_instance, repo_id, "security_token")
	if status.Code(err) == codes.NotFound {
		return "", nil
	}
	return token, err
}
