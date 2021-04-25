package discord_structs

// Message A struct outlining the content and destination of a message
type Message struct {
	ChannelID string
	Message string
	Mentions []string
}