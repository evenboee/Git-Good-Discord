package discord

import (
	"errors"
	"git-good-discord/discord/discord_structs"
)

func (i Implementation) SendMessage(msg discord_structs.Message) error {
	if session == nil { return errors.New("session is not open") }

	_, err := session.ChannelMessageSend(msg.ChannelID, buildMentions(msg.Mentions) + msg.Message)
	return err
}

func buildMentions(mentions []string) string {
	mentionString := ""
	for _, mention := range mentions {
		mentionString += mention
	}
	return mentionString
}