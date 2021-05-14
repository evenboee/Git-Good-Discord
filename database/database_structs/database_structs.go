package database_structs

// Subscriber describes a discord subscriber
type Subscriber struct {

	// DiscordUserId is the unique discord ID
	DiscordUserId string `json:"discord_user_id" firestore:"-"`

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
