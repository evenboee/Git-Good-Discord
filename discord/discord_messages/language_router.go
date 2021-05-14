package discord_messages

import (
	"git-good-discord/database/database_structs"
	"git-good-discord/discord/discord_structs"
	"git-good-discord/gitlab/gitlab_structs"
	"git-good-discord/utils"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
	"time"
)

// GetPing gets the message for the ping command
func GetPing(language string, info []string, session *discordgo.Session, messageCreate *discordgo.MessageCreate) discord_structs.Message {
	pingLanguage := getLanguage(language).Ping

	response := ""
	mentions := make([]string, 0)

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

func GetSetAccessToken(command string, language string, newAccessToken string, expectedParts string, messageCreate *discordgo.MessageCreate) discord_structs.EmbeddedMessage {
	languagePack := getLanguage(language).SetAccessToken
	response := ""
	switch command {
	case "WrongParts":
		response = placeholderHandler(languagePack.WrongParts, expectedParts)
	case "WrongPath":
		response = languagePack.WrongPath
	case "PathElementEmpty":
		response = placeholderHandler(languagePack.PathElementEmpty, expectedParts)
	case "AddTokenFail":
		response = languagePack.AddTokenFail
	case "Successful":
		response = placeholderHandler(languagePack.Successful, newAccessToken)
	}

	return discord_structs.EmbeddedMessage{
		Message:      discord_structs.Message{
			ChannelID: messageCreate.ChannelID,
			Message:   response,
			Mentions:  []string{messageCreate.Author.Mention()},
		},
		MessageEmbed: discordgo.MessageEmbed{},
	}
}

// GetReloadLanguage gets the message for the Reload Language command
func GetReloadLanguage(language string, action string, messageCreate *discordgo.MessageCreate) discord_structs.EmbeddedMessage {
	reloadLanguage := getLanguage(language).ReloadLanguage

	response := reloadLanguage.SuccessfullyReloaded
	if action == "errorReloading" {
		response = reloadLanguage.ErrorReloading
	}

	return discord_structs.EmbeddedMessage{
		Message: discord_structs.Message{
			ChannelID: messageCreate.ChannelID,
			Message:   response,
			Mentions:  []string{messageCreate.Author.Mention()},
		},
	}
}

// GetChangeLanguage gets the message for the Change Language command
func GetChangeLanguage(command string, language string, newLanguage string, messageCreate *discordgo.MessageCreate) discord_structs.EmbeddedMessage {
	languagePack := getLanguage(language).ChangeLanguage
	response := ""
	switch command {
	case "NoParam":
		response = languagePack.NoParam
	case "Invalid":
		response = placeholderHandler(languagePack.InvalidLanguage, newLanguage)
	case "DatabaseSetFail":
		response = languagePack.DatabaseSetFail
	case "Success":
		response = placeholderHandler(languagePack.Successful, newLanguage)
	}

	return discord_structs.EmbeddedMessage{
		Message: discord_structs.Message{
			ChannelID: messageCreate.ChannelID,
			Message:   response,
			Mentions:  []string{messageCreate.Author.Mention()},
		},
	}
}

// NotAuthorizedMessage gets the message for when you are not authorized
func NotAuthorizedMessage(language string, command string, messageCreate *discordgo.MessageCreate) discord_structs.EmbeddedMessage {
	response := "You are not authorized to do this!"
	if command == "ChangeLanguage" {
		response = getLanguage(language).ChangeLanguage.NotAuthorized
	} else if command == "SetPrefix" {
		response = getLanguage(language).SetLanguagePrefix.NotAuthorized
	} else if command == "SetAccessToken" {
		response = getLanguage(language).SetAccessToken.NotAuthorized
	}

	return discord_structs.EmbeddedMessage{
		Message: discord_structs.Message{
			ChannelID: messageCreate.ChannelID,
			Message:   response,
			Mentions:  []string{messageCreate.Author.Mention()},
		},
	}
}

// GetCommandNotRecognized gets the message for when a command is not recognized
func GetCommandNotRecognized(language string, m *discordgo.MessageCreate) discord_structs.EmbeddedMessage{
	return discord_structs.EmbeddedMessage{
		Message: discord_structs.Message{
			ChannelID: m.ChannelID,
			Message:   getLanguage(language).Errors.CommandNotRecognized,
			Mentions:  []string{m.Author.Mention()},
		},
	}
}

// SetPrefix gets the message for the Set Prefix command
func SetPrefix(command string, prefix string, language string, m *discordgo.MessageCreate) discord_structs.EmbeddedMessage {
	languagePack := getLanguage(language).SetLanguagePrefix
	response := ""
	if command == "dbError" {
		response = languagePack.NotAuthorized
	} else {
		response = placeholderHandler(languagePack.Successful, prefix)
	}

	return discord_structs.EmbeddedMessage{
		Message: discord_structs.Message{
			ChannelID: m.ChannelID,
			Message:   response,
			Mentions:  []string{m.Author.Mention()},
		},
	}
}

// GetHelp gets the message for the help command
func GetHelp(prefix string, language string, isAdmin bool, messageCreate *discordgo.MessageCreate) discord_structs.EmbeddedMessage {
	helpLanguage := getLanguage(language).HelpCommand
	response := "\n***Commands***\n> " +
		prefix + "help - " + helpLanguage.Help + "\n> " +
		prefix + "subscribe <instance>/<repo_id>/<gitlab_username> <type1,type2,...> - " + helpLanguage.Subscribe + "\n> " +
		prefix + "unsubscribe <instance>/<repo_id>/<gitlab_username> - " + helpLanguage.Unsubscribe + "\n> " +
		prefix + "get <channel-name> - " + helpLanguage.Get + "\n> " +
		prefix + "ping <group> - " + helpLanguage.Ping + "\n"

	if isAdmin {
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

// GetSubscribe gets the message for the subscribe command
func GetSubscribe(command string, variable string, variable2 string, language string, m *discordgo.MessageCreate) discord_structs.EmbeddedMessage {
	languagePack := getLanguage(language).Subscribe
	response := ""
	switch command {
	case "PathFormatError":
		response = languagePack.PathFormatError
	case "TokenGenerationFail":
		response = languagePack.TokenGenerationFail
	case "InvocationURLFail":
		response = languagePack.InvocationURLFail
	case "RepoIDFormatError":
		response = languagePack.RepoIDFormatError
	case "AccessTokenFail":
		response = languagePack.AccessTokenFail
	case "WebhookRegistrationError":
		response = languagePack.WebhookRegistrationError
	case "DatabaseAddSecurityTokenFail":
		response = languagePack.DatabaseAddSecurityTokenFail
	case "DatabaseAddFail":
		response = languagePack.DatabaseAddFail
	case "Successful":
		response = placeholderHandler(languagePack.Successful, variable, variable2)
	}
	return discord_structs.EmbeddedMessage{
		Message: discord_structs.Message{
			ChannelID: m.ChannelID,
			Message:   response,
			Mentions:  []string{m.Author.Mention()},
		},
	}
}

// GetUnsubscribe gets the message for the unsubscribe command
func GetUnsubscribe(language string, command string, variable string, m *discordgo.MessageCreate) discord_structs.EmbeddedMessage {
	languagePack := getLanguage(language).Unsubscribe
	response := ""

	switch command {
	case "PartsError":
		response = placeholderHandler(languagePack.PartsError, variable)
	case "PathFormatError":
		response = languagePack.PathFormatError
	case "DatabaseRemoveFail":
		response = languagePack.DatabaseRemoveFail
	case "Successful":
		response = placeholderHandler(languagePack.Successful, variable)
	}

	return discord_structs.EmbeddedMessage{
		Message: discord_structs.Message{
			ChannelID: m.ChannelID,
			Message:   response,
			Mentions:  []string{m.Author.Mention()},
		},
	}
}

// NotifySubscribers gets the message for subscribers when a gitlab event has
// occured
func NotifySubscribers(language string, discordChannelID string, subscribers []database_structs.Subscriber, notification gitlab_structs.WebhookNotification) discord_structs.EmbeddedMessage {
	mentions := make([]string, 1)
	for _, subscriber := range subscribers {
		discordUser := &discordgo.User{ID: subscriber.DiscordUserId}
		mentions = append(mentions, discordMention(discordUser))
	}

	authorURL := utils.HTTPS(notification.Project.URL + "/" + notification.User.Username)

	timeStamp, err := time.Parse("2006-01-02T15:04:05Z", strings.ReplaceAll(strings.ReplaceAll(notification.ObjectAttributes.CreatedAt, " ", "T"), "TUTC", "Z"))

	if err != nil {
		timeStamp = time.Time{}
	}

	return discord_structs.EmbeddedMessage{
		Message: discord_structs.Message{
			ChannelID: discordChannelID,
			Message:   getWebhookNotificationMessage(language, notification),
			Mentions:  mentions,
		},

		MessageEmbed: discordgo.MessageEmbed{
			URL:         notification.ObjectAttributes.URL,
			Type:        discordgo.EmbedTypeLink,
			Title:       notification.ObjectAttributes.Title,
			Description: notification.ObjectAttributes.Description,
			Timestamp:   timeStamp.Format("2006-01-02T15:04:05-0700"),
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

// getWebhookNotificationMessage gets the correct message based on the type of
// gitlab even that occured
func getWebhookNotificationMessage(language string, notification gitlab_structs.WebhookNotification) string {
	switch notification.ObjectKind {
	case gitlab_structs.NotificationMergeRequest:
		return placeholderHandler(getLanguage(language).NotificationMergeRequest.Success, notification.User.Name)
	case gitlab_structs.NotificationIssue:
		return placeholderHandler(getLanguage(language).NotificationIssue.Success, notification.User.Name)
	}

	log.Printf("Unexpected notification type '%s'", notification.ObjectKind)
	return "Oopsie, Something went wrong!"
}
