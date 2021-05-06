package discord_messages

import (
	"git-good-discord/discord/discord_structs"
	"github.com/bwmarrin/discordgo"
	"strings"
)

var currentLanguagePack commands

func GetChannel(messageCreate *discordgo.MessageCreate, prefix string) discord_structs.Message {
	getChannelLanguage := currentLanguagePack.GetChannel

	_, info := splitMessage(messageCreate.Content, prefix)
	response := ""
	if len(info) == 0 {
		response = getChannelLanguage.NotSpecified
	} else {
		if info[0] == "channel_id" {
			response = "Channel_id: " + messageCreate.ChannelID
		} else {
			response = placeholderHandler(getChannelLanguage.NotRecognized, info[0])
		}
	}

	return discord_structs.Message{
		ChannelID: messageCreate.ChannelID,
		Message:   response,
		Mentions:  nil,
	}
}

func Ping(session *discordgo.Session, messageCreate *discordgo.MessageCreate, prefix string) discord_structs.Message {
	pingLanguage := currentLanguagePack.Ping

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
		response = pingLanguage.ErrorGettingRoles
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
			response = placeholderHandler(pingLanguage.RoleNotFound, role)
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