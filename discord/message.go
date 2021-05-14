package discord

import (
	"errors"
	"git-good-discord/discord/discord_structs"
	"github.com/bwmarrin/discordgo"
	"reflect"
)

func (i Implementation) sendMessage(msg discord_structs.Message) error {
	_, err := session.ChannelMessageSend(msg.ChannelID, buildMentions(msg.Mentions)+msg.Message)
	return err
}

func (i Implementation) SendMessage(msg discord_structs.EmbeddedMessage) error {
	if session == nil {
		return errors.New("session is not open")
	}
	var err error
	if reflect.DeepEqual(msg.MessageEmbed, discordgo.MessageEmbed{}) {
		err = i.sendMessage(msg.Message)
	} else {
		_, err = session.ChannelMessageSendComplex(msg.Message.ChannelID, &discordgo.MessageSend{
			Content: buildMentions(msg.Message.Mentions) + msg.Message.Message,
			Embed:   &msg.MessageEmbed,
		})
	}
	return err
}

// buildMentions builds a mentions-string
func buildMentions(mentions []string) string {
	mentionString := ""
	for _, mention := range mentions {
		mentionString += mention
	}
	return mentionString + " "
}
