package database_structs

// Subscriber describes a discord subscriber
type Subscriber struct {
	// DiscordUserId is the unique discord ID
	DiscordUserId string `json:"id" firestore:"id"`

	// ChannelID is the channel the subscription is made to
	ChannelID       string `json:"channel_id" firestore:"channel_id"`

	// Instance, RepoID and GitlabUsername are to specify path of a subscription
	Instance        string `json:"instance" firestore:"instance"`
	RepoID          string `json:"repo_id" firestore:"repo_id"`
	GitlabUsername  string  `json:"gitlab_user" firestore:"gitlab_user"`

	// Issues defines if this subscriber is interested in issues
	Issues        bool   `json:"issues" firestore:"issues"`

	// MergeRequests defines if this subscriber is interested in merge requests
	MergeRequests bool   `json:"merge_requests" firestore:"merge_requests"`
}

// ChannelSettings describes settings for a given discord channel
type ChannelSettings struct {

	// Language is the language for the channel
	Language string `json:"language" firestore:"language"`

	// Prefix is the command prefix for the channel
	Prefix   string `json:"prefix" firestore:"prefix"`

}
