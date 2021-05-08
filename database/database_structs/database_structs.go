package database_structs

type Subscriber struct {
	DiscordUserId string `json:"discord_user_id" firestore:"-"`
	Issues        bool   `json:"issues" firestore:"issues"`
	MergeRequests bool   `json:"merge_requests" firestore:"merge_requests"`
}
