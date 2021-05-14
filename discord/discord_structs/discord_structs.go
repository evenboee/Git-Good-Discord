package discord_structs

import "github.com/bwmarrin/discordgo"

// Message A struct outlining the content and destination of a message
type Message struct {

	// ChannelID is the discord channel ID
	ChannelID string

	// Message is the message
	Message string

	// Mentions is the usernames that should be mentioned
	Mentions []string
}

// EmbeddedMessage is a type of message that can have embedded content
type EmbeddedMessage struct {

	// Message without embedded content
	Message

	// Embedded content
	discordgo.MessageEmbed
}
