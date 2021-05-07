package discord_structs

import "github.com/bwmarrin/discordgo"

// Message A struct outlining the content and destination of a message
type Message struct {
	ChannelID string
	Message string
	Mentions []string
}

type EmbeddedMessage struct {
	Message
	discordgo.MessageEmbed
}