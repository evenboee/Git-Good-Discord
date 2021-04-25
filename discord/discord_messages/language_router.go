package discord_messages

import (
	"git-good-discord/discord/discord_structs"
	"github.com/bwmarrin/discordgo"
	"strings"
)

var language = Norwegian{}

type languageCommands interface {
	getChannel() []string
	ping() []string
}

type English struct{}
type Norwegian struct{}

func getCurrentLanguage() languageCommands {
	return language
}

func GetChannel(messageCreate *discordgo.MessageCreate, prefix string) discord_structs.Message {
	languageStrings := getCurrentLanguage().getChannel()
	_, info := splitMessage(messageCreate.Content, prefix)
	response := ""
	if len(info) == 0 {
		response = languageStrings[0]
	} else {
		if info[0] == "channel_id" {
			response = "Channel_id: " + messageCreate.ChannelID
		} else {
			response = languageStrings[1] + info[0]
		}
	}

	return discord_structs.Message{
		ChannelID: messageCreate.ChannelID,
		Message:   response,
		Mentions:  nil,
	}
}

func Ping(session *discordgo.Session, messageCreate *discordgo.MessageCreate, prefix string) discord_structs.Message {
	languageStrings := getCurrentLanguage().ping()

	response := ""
	mentions := make([]string, 0)
	_, info := splitMessage(messageCreate.Content, prefix)

	role := ""
	if len(info) == 0 {
		role = "everyone"
	} else {
		role = info[0]
	}
	role = strings.ToLower(role)
	roles, err := session.GuildRoles(messageCreate.GuildID)
	if err != nil {
		response = languageStrings[0]
	}
	for _, r := range roles {
		if strings.Contains(strings.ToLower(r.Name), role) {
			mention := discordMention(r)
			if mention != "" {
				mentions = append(mentions, mention)
			}
		}
	}
	if response == "" {
		if len(mentions) == 0 {
			response = languageStrings[1] + " \"" + role + "\""
		} else {
			if len(info) >= 2 {
				response = " " + strings.Join(info[1:], " ")
			} else {
				response = " ping"
			}
		}
	}
	return discord_structs.Message{
		ChannelID: messageCreate.ChannelID,
		Message:   response,
		Mentions:  mentions,
	}
}
