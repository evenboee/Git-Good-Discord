package discord

import (
	"fmt"
	"git-good-discord/database/database_interfaces"
	"git-good-discord/discord/discord_messages"
	"git-good-discord/discord/discord_structs"
	"git-good-discord/gitlab/gitlab_interfaces"
	"git-good-discord/gitlab/gitlab_structs"
	"git-good-discord/utils"
	"github.com/bwmarrin/discordgo"
	"log"
	"strconv"
	"strings"
)

// commandHandler handles commands
func commandHandler(i Implementation, s *discordgo.Session, m *discordgo.MessageCreate, prefix string, language string) {
	parts := strings.Split(m.Content, " ")
	command := strings.Trim(strings.ToLower(parts[0]), prefix)

	var message discord_structs.EmbeddedMessage
	switch strings.ToLower(command) {
	case "get":
		message = getHandler(prefix, m)
	case "access_token":
		if !memberIsAdmin(m, s) {
			message = discord_messages.NotAuthorizedMessage(language, "SetAccessToken", m)
		} else {
			message = getSetAccessToken(language, i.DatabaseService, s, m)
		}
	case "ping":
		_, info := splitMessage(m.Content, prefix)
		message = discord_structs.EmbeddedMessage{Message: discord_messages.GetPing(language, info, s, m)}
	case "subscribe":
		message = subscribeHandler(language, i.DatabaseService, i.GitlabService, m)
	case "unsubscribe":
		message = unsubscribeHandler(language, i.DatabaseService, m)
	case "subscriptions":
		message = getSubscriptions(language, i.DatabaseService, m)
	case "reload":
		if !memberIsAdmin(m, s) {
			message = discord_messages.NotAuthorizedMessage(language, "ReloadLanguage", m)
		} else {
			message = reloadHandler(language, m)
		}
	case "language":
		if !memberIsAdmin(m, s) {
			message = discord_messages.NotAuthorizedMessage(language, "ChangeLanguage", m)
		} else {
			message = languageHandler(prefix, language, i.DatabaseService, m)
		}
	case "help":
		message = discord_messages.GetHelp(prefix, language, memberIsAdmin(m, s), m)
	default:
		message = discord_messages.GetCommandNotRecognized(language, m)
	}
	err := i.SendMessage(message)
	if err != nil {
		log.Printf("Bot.go: %v\n", err)
		return
	}
}

func getSubscriptions(language string, db database_interfaces.Database, m *discordgo.MessageCreate) discord_structs.EmbeddedMessage {
	subscriptions, err := db.GetConnection().GetAllSubscriptions(m.ChannelID, m.Author.ID)
	command := ""
	message := ""
	if err == nil {
		for _, subscription := range subscriptions {
			var events []string
			if subscription.MergeRequests {
				events = append(events, "merge_requests")
			}
			if subscription.Issues {
				events = append(events, "issues")
			}
			message += fmt.Sprintf("- %s/%s/%s %s\n", subscription.Instance, subscription.RepoID, subscription.GitlabUsername, strings.Join(events, ","))
		}
		command = "Successful"
	} else {
		command = "DatabaseFail"
	}

	return discord_messages.GetSubscriptions(command, language, message, m)
}

func getSetAccessToken(language string, db database_interfaces.Database, s *discordgo.Session, m *discordgo.MessageCreate) discord_structs.EmbeddedMessage {
	parts := strings.Split(m.Content, " ")

	// !access_token {path}
	if len(parts) != 2 {
		return discord_messages.GetSetAccessToken("WrongParts", language, "", fmt.Sprintf("%d", len(parts)), m)
	}

	path := strings.Split(parts[1], "/")
	// instance/project_id/token
	if len(path) != 3 {
		return discord_messages.GetSetAccessToken("WrongPath", language, "", "", m)
	}

	for i, v := range path {
		if v == "" {
			return discord_messages.GetSetAccessToken("PathElementEmpty", language, "", fmt.Sprintf("%d", i+1), m)
		}
	}

	err := db.GetConnection().AddAccessToken(m.ChannelID, path[0], path[1], path[2])
	if err != nil {
		return discord_messages.GetSetAccessToken("AddTokenFail", language, "", "", m)
	}

	return discord_messages.GetSetAccessToken("Successful", language, path[2], "", m)
}

// reloadHandler handles reload command
func reloadHandler(language string, messageCreate *discordgo.MessageCreate) discord_structs.EmbeddedMessage {
	err := discord_messages.ReloadLanguageFiles()
	action := ""
	if err != nil {
		log.Print("Problem reloading language pack: ")
		log.Println(err)
		action = "errorReloading"
	}

	return discord_messages.GetReloadLanguage(language, action, messageCreate)
}

// languageHandler handles language command
func languageHandler(prefix string, language string, db database_interfaces.Database, messageCreate *discordgo.MessageCreate) discord_structs.EmbeddedMessage {
	_, info := splitMessage(messageCreate.Content, prefix)
	command := ""
	newLanguage := ""
	if len(info) == 0 {
		command = "NoParam"
	} else {
		newLanguage = strings.ToLower(info[0])
		if !discord_messages.IsLanguage(newLanguage) {
			//Language is not available
			command = "Invalid"
		} else {
			command = "Success"

			//TODO: Consider moving
			err := db.GetConnection().SetChannelLanguage(messageCreate.ChannelID, newLanguage)
			if err != nil {
				log.Printf("Language handler - %v\n", err)
				command = "DatabaseSetFail"
			}
		}
	}
	return discord_messages.GetChangeLanguage(command, language, newLanguage, messageCreate)
}

// setPrefixHandler handles the get prefix command
func setPrefixHandler(newPrefix string, language string, db database_interfaces.Database, m *discordgo.MessageCreate) discord_structs.EmbeddedMessage {
	//TODO: Consider moving
	err := db.GetConnection().SetChannelPrefix(m.ChannelID, newPrefix)
	command := "Success"
	if err != nil {
		log.Printf("SetPrefix: %v\n", err)
		command = "dbError"
	}
	return discord_messages.SetPrefix(command, newPrefix, language, m)
}

// unsubscribeHandler handles the unsubscribe command
func unsubscribeHandler(language string, db database_interfaces.Database, m *discordgo.MessageCreate) discord_structs.EmbeddedMessage {
	command := ""
	variable := ""

	parts := strings.Split(m.Content, " ")
	if len(parts) != 2 {
		command = "PartsError"
		variable = fmt.Sprintf("%d", len(parts))
	} else {
		path := strings.Split(parts[1], "/")
		if len(path) != 3 {
			command = "PathFormatError"
		} else {
			instance := path[0]
			repoID := path[1]
			gitlabUsername := path[2]
			err := db.GetConnection().DeleteSubscriber(m.ChannelID, instance, repoID, gitlabUsername, m.Author.ID)
			if err != nil {
				command = "DatabaseRemoveFail"
			} else {
				command = "Successful"
				variable = parts[1]
			}
		}
	}
	return discord_messages.GetUnsubscribe(language, command, variable, m)
}

// subscribeHandler handles the subscribe command
func subscribeHandler(language string, db database_interfaces.Database, gitlab gitlab_interfaces.Interface, m *discordgo.MessageCreate) discord_structs.EmbeddedMessage {
	command, variable, variable2 := subscribeGetCommands(db, gitlab, m)
	return discord_messages.GetSubscribe(command, variable, variable2, language, m)
}

func subscribeGetCommands(db database_interfaces.Database, gitlab gitlab_interfaces.Interface, m *discordgo.MessageCreate) (string, string, string) {
	//TODO: Should consider moving this to abstraction
	parts := strings.Split(m.Content, " ")
	command := ""
	variable := ""
	variable2 := ""
	if len(parts) != 3 {
		command = "PathFormatError"
		return command, variable, variable2
	}
	path := strings.Split(parts[1], "/")

	if len(path) != 3 {
		command = "PathFormatError"
		return command, variable, variable2
	}
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

	if len(newSubscriptions) == 0 {
		command = "PathFormatError"
		return command, variable, variable2
	}

	token, err := utils.GenerateUUID()
	if err != nil {
		command = "TokenGenerationFail"
		return command, variable, variable2
	}

	url, err := gitlab.GetWebhookInvocationURL(m.ChannelID)
	if err != nil {
		command = "InvocationURLFail"
		return command, variable, variable2
	}

	id, err := strconv.Atoi(repoID)
	if err != nil {
		command = "RepoIDFormatError"
		return command, variable, variable2
	}

	accessToken, err := db.GetConnection().GetAccessToken(m.ChannelID, instance, repoID)
	if err != nil {
		command = "AccessTokenFail"
		return command, variable, variable2
	}

	project := gitlab_structs.Project{
		URL:         instance,
		ID:          id,
		AccessToken: accessToken,
	}
	webhook := gitlab_structs.Webhook{
		URL:                 url,
		SecretToken:         token,
		IssuesEvents:        true,
		MergeRequestsEvents: true,
	}

	_, err = gitlab.RegisterWebhook(project, webhook)
	if err != nil {
		if !strings.Contains(err.Error(), "webhook is already registered") {
			command = "WebhookRegistrationError"
			return command, variable, variable2
		}
	} else {
		err = db.GetConnection().AddSecurityToken(m.ChannelID, instance, repoID, token)
		if err != nil {
			command = "DatabaseAddSecurityTokenFail"
			return command, variable, variable2
		}
	}

	err = db.GetConnection().AddSubscriber(m.ChannelID, instance, repoID, gitlabUsername, m.Author.ID, issues, merge_requests)
	if err != nil {
		command = "DatabaseAddFail"
	} else {
		command = "Successful"
		variable = parts[1]
		variable2 = strings.Join(newSubscriptions, ", ")
	}
	return command, variable, variable2
}

//Testing functions
func getHandler(prefix string, messageCreate *discordgo.MessageCreate) discord_structs.EmbeddedMessage {
	_, info := splitMessage(messageCreate.Content, prefix)
	response := "Specify what you want to get"
	if len(info) != 0 {
		if info[0] == "channel_id" {
			response = "Channel_id: " + messageCreate.ChannelID
		} else {
			response = "Could not recognize " + info[0]
		}
	}

	return discord_structs.EmbeddedMessage{
		Message: discord_structs.Message{
			ChannelID: messageCreate.ChannelID,
			Message:   response,
			Mentions:  nil,
		},
	}
}
