package discord_messages

import (
	"git-good-discord/database/database_interfaces"
	"git-good-discord/discord/discord_structs"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

var currentLanguagePack commands

func GetGetChannel(messageCreate *discordgo.MessageCreate, prefix string, language string) discord_structs.Message {
	var getChannelLanguage GetChannel
	if v, ok := languageFiles[language]; ok {
		getChannelLanguage = v.GetChannel
	} else {
		getChannelLanguage = languageFiles["english"].GetChannel
	}

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

func GetPing(session *discordgo.Session, messageCreate *discordgo.MessageCreate, prefix string, language string) discord_structs.Message {
	var pingLanguage Ping
	if v, ok := languageFiles[language]; ok {
		pingLanguage = v.Ping
	} else {
		pingLanguage = languageFiles["english"].Ping
	}

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

func GetReloadLanguage(messageCreate *discordgo.MessageCreate, language string) discord_structs.EmbeddedMessage{
	var reloadLanguage ReloadLang
	if v, ok := languageFiles[language]; ok {
		reloadLanguage = v.ReloadLanguage
	} else {
		reloadLanguage = languageFiles["english"].ReloadLanguage
	}

	response := ""

	err := ReloadLanguageFiles()
	if err != nil {
		log.Print("Problem reloading language pack: ")
		log.Println(err)
		response = reloadLanguage.ErrorReloading
	} else {
		response = reloadLanguage.SuccessfullyReloaded
		currentLanguagePack = languageFiles[currentLanguagePack.Language]
	}


	return discord_structs.EmbeddedMessage{
		Message:      discord_structs.Message{
			ChannelID: messageCreate.ChannelID,
			Message:   response,
			Mentions: []string{messageCreate.Author.Mention()},
		},
	}
}

func GetChangeLanguage(db database_interfaces.Database, s *discordgo.Session, messageCreate *discordgo.MessageCreate, prefix string, language string) discord_structs.EmbeddedMessage{
	var changeLanguage ChangeLanguage
	if v, ok := languageFiles[language]; ok {
		changeLanguage = v.ChangeLanguage
	} else {
		changeLanguage = languageFiles["english"].ChangeLanguage
	}

	roles, err := s.GuildRoles(messageCreate.GuildID)
	if err != nil { return discord_structs.EmbeddedMessage{} }

	isAdmin := memberIsAdmin(messageCreate.Member, roles)

	if !isAdmin {
		return discord_structs.EmbeddedMessage{
			Message: discord_structs.Message{
				ChannelID: messageCreate.ChannelID,
				Message:   changeLanguage.NotAuthorized,
				Mentions:  []string{messageCreate.Author.Mention()},
			},
		}
	}

	response := ""
	_, info := splitMessage(messageCreate.Content, prefix)
	if len(info) == 0 {
		response = changeLanguage.NoParam
	} else {
		nLanguage := strings.ToLower(info[0])
		if languageFiles[nLanguage] == (commands{}) {
			//Language is not available
			response = placeholderHandler(changeLanguage.InvalidLanguage, nLanguage)
		} else {
			response = placeholderHandler(changeLanguage.Successful, nLanguage)
			currentLanguagePack = languageFiles[nLanguage]

			err := db.GetConnection().SetChannelLanguage(messageCreate.ChannelID, nLanguage)
			if err != nil {
				response = changeLanguage.DatabaseSetFail
			}
		}
	}

	return discord_structs.EmbeddedMessage{
		Message:      discord_structs.Message{
			ChannelID: messageCreate.ChannelID,
			Message:   response,
			Mentions: []string{messageCreate.Author.Mention()},
		},
	}
}

func SetPrefix(db database_interfaces.Database, s *discordgo.Session, m *discordgo.MessageCreate, nPrefix string, language string) discord_structs.Message {
	var languagePack setPrefix
	if v, ok := languageFiles[language]; ok {
		languagePack = v.SetPrefix
	} else {
		languagePack = languageFiles["english"].SetPrefix
	}

	roles, err := s.GuildRoles(m.GuildID)
	if err != nil { return discord_structs.Message{} }

	isAdmin := memberIsAdmin(m.Member, roles)

	if !isAdmin {
		return discord_structs.Message{
			ChannelID: m.ChannelID,
			Message:   languagePack.NotAuthorized,
			Mentions:  []string{m.Author.Mention()},
		}
	}

	err = db.GetConnection().SetChannelPrefix(m.ChannelID, nPrefix)
	if err != nil {
		return discord_structs.Message{
			ChannelID: m.ChannelID,
			Message:   languagePack.NotAuthorized,
			Mentions:  []string{m.Author.Mention()},
		}
	}

	return discord_structs.Message{
		ChannelID: m.ChannelID,
		Message:   placeholderHandler(languagePack.Successful, nPrefix),
		Mentions:  []string{m.Author.Mention()},
	}
}