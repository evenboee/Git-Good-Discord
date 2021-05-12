package discord_messages

import (
	"fmt"
	"git-good-discord/database/database_interfaces"
	"git-good-discord/database/database_structs"
	"git-good-discord/discord/discord_structs"
	"git-good-discord/gitlab/gitlab_structs"
	"git-good-discord/utils"
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

func GetReloadLanguage(messageCreate *discordgo.MessageCreate, language string) discord_structs.EmbeddedMessage {
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
		Message: discord_structs.Message{
			ChannelID: messageCreate.ChannelID,
			Message:   response,
			Mentions:  []string{messageCreate.Author.Mention()},
		},
	}
}

func GetChangeLanguage(db database_interfaces.Database, s *discordgo.Session, messageCreate *discordgo.MessageCreate, prefix string, language string) discord_structs.EmbeddedMessage {
	var changeLanguage ChangeLanguage
	if v, ok := languageFiles[language]; ok {
		changeLanguage = v.ChangeLanguage
	} else {
		changeLanguage = languageFiles["english"].ChangeLanguage
	}

	roles, err := s.GuildRoles(messageCreate.GuildID)
	if err != nil {
		return discord_structs.EmbeddedMessage{}
	}

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
		Message: discord_structs.Message{
			ChannelID: messageCreate.ChannelID,
			Message:   response,
			Mentions:  []string{messageCreate.Author.Mention()},
		},
	}
}

func SetPrefix(db database_interfaces.Database, s *discordgo.Session, m *discordgo.MessageCreate, nPrefix string, language string) discord_structs.Message {
	var languagePack SetLanguagePrefix
	if v, ok := languageFiles[language]; ok {
		languagePack = v.SetLanguagePrefix
	} else {
		languagePack = languageFiles["english"].SetLanguagePrefix
	}

	roles, err := s.GuildRoles(m.GuildID)
	if err != nil {
		return discord_structs.Message{}
	}

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

func GetHelp(s *discordgo.Session, messageCreate *discordgo.MessageCreate, prefix string) discord_structs.EmbeddedMessage {
	helpLanguage := currentLanguagePack.HelpCommand

	roles, err := s.GuildRoles(messageCreate.GuildID)
	if err != nil {
		return discord_structs.EmbeddedMessage{}
	}
	response := "\n***Commands***\n> " +
		prefix + "help - " + helpLanguage.Help + "\n> " +
		prefix + "subscribe <instance>/<repo_id>/<gitlab_username> <type1,type2,...> - " + helpLanguage.Subscribe + "\n> " +
		prefix + "unsubscribe <instance>/<repo_id>/<gitlab_username> - " + helpLanguage.Unsubscribe + "\n> " +
		prefix + "get <channel-name> - " + helpLanguage.Get + "\n> " +
		prefix + "ping <group> - " + helpLanguage.Ping + "\n"

	if memberIsAdmin(messageCreate.Member, roles) {
		response += "\n" +
			"***Admin commands***\n> " +
			prefix + "reload - " + helpLanguage.Reload + " " + helpLanguage.AdminOnly + "\n> " +
			prefix + "language <language> - " + helpLanguage.Reload + " " + helpLanguage.AdminOnly + "\n> " +
			"!" + "set prefix <prefix> - " + helpLanguage.SetPrefix + " " + helpLanguage.AdminOnly + "\n"
	}

	return discord_structs.EmbeddedMessage{
		Message: discord_structs.Message{
			ChannelID: messageCreate.ChannelID,
			Message:   response,
			Mentions:  []string{messageCreate.Author.Mention()},
		},
	}
}

func GetSubscribe(db database_interfaces.Database, m *discordgo.MessageCreate, language string) discord_structs.Message {
	response := ""

	var languagePack Subscribe
	if v, ok := languageFiles[language]; ok {
		languagePack = v.Subscribe
	} else {
		languagePack = languageFiles["english"].Subscribe
	}

	parts := strings.Split(m.Content, " ")
	if len(parts) == 3 {
		path := strings.Split(parts[1], "/")
		if len(path) == 3 {
			instance := path[0]
			repoID := path[1]
			gitlabUsername := path[2]

			issues := false
			merge_requests := false

			subscriptions := strings.Split(parts[2], ",")
			var newSubscriptions []string
			for _, v := range subscriptions {
				switch v {
				case "issues":
					issues = true
					newSubscriptions = append(newSubscriptions, "issues")
				case "merge_requests":
					merge_requests = true
					newSubscriptions = append(newSubscriptions, "merge_requests")
				}
			}

			err := db.GetConnection().AddSubscriber(m.ChannelID, instance, repoID, gitlabUsername, m.Author.ID, issues, merge_requests)
			if err != nil {
				response = languagePack.DatabaseAddFail
			} else {
				response = placeholderHandler(languagePack.Successful, parts[1], strings.Join(newSubscriptions, ","))
			}
		} else {
			response = languagePack.PathFormatError
		}
	} else {
		response = placeholderHandler(languagePack.PathFormatError, fmt.Sprintf("%d", len(parts)))
	}

	return discord_structs.Message{
		ChannelID: m.ChannelID,
		Message:   response,
		Mentions:  []string{m.Author.Mention()},
	}
}

func GetUnsubscribe(db database_interfaces.Database, m *discordgo.MessageCreate, language string) discord_structs.Message {
	// Unsubscribe = delete subscriber
	response := ""

	var languagePack Unsubscribe
	if v, ok := languageFiles[language]; ok {
		languagePack = v.Unsubscribe
	} else {
		languagePack = languageFiles["english"].Unsubscribe
	}

	parts := strings.Split(m.Content, " ")
	if len(parts) == 2 {
		path := strings.Split(parts[1], "/")
		if len(path) == 3 {
			instance := path[0]
			repoID := path[1]
			gitlabUsername := path[2]
			err := db.GetConnection().DeleteSubscriber(m.ChannelID, instance, repoID, gitlabUsername, m.Author.ID)
			if err != nil {
				response = languagePack.DatabaseRemoveFail
			}
			response = placeholderHandler(languagePack.Successful, parts[1])
		} else {
			response = languagePack.PathFormatError
		}
	} else {
		response = placeholderHandler(languagePack.PartsError, fmt.Sprintf("%d", len(parts)))
	}

	return discord_structs.Message{
		ChannelID: m.ChannelID,
		Message:   response,
		Mentions:  []string{m.Author.Mention()},
	}
}

func NotifySubscribersOfMergeRequest(discordChannelID string, subscribers []database_structs.Subscriber, notification gitlab_structs.MergeRequestWebhookNotification) discord_structs.EmbeddedMessage {
	notificationMergeRequestLanguage := currentLanguagePack.NotificationMergeRequest
	mentions := make([]string, 1, 1)
	for _, subscriber := range subscribers {
		discordUser := &discordgo.User{ID: subscriber.DiscordUserId}
		mentions = append(mentions, discordMention(discordUser))
	}

	authorURL := utils.HTTPS(notification.Project.URL + "/" + notification.User.Username)

	return discord_structs.EmbeddedMessage{
		Message: discord_structs.Message{
			ChannelID: discordChannelID,
			Message:   placeholderHandler(notificationMergeRequestLanguage.Success, notification.User.Name),
			Mentions:  mentions,
		},
		MessageEmbed: discordgo.MessageEmbed{
			URL:         notification.ObjectAttributes.URL,
			Type:        discordgo.EmbedTypeLink,
			Title:       notification.ObjectAttributes.Title,
			Description: notification.ObjectAttributes.Description,
			Timestamp:   notification.ObjectAttributes.CreatedAt,
			Author: &discordgo.MessageEmbedAuthor{
				URL:  authorURL,
				Name: notification.User.Name,
			},
			Provider: &discordgo.MessageEmbedProvider{
				URL:  notification.ObjectAttributes.URL,
				Name: "Gitlab",
			},
		},
	}
}
