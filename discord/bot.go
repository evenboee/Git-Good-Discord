package discord

import (
	"fmt"
	"git-good-discord/discord/discord_messages"
	"git-good-discord/discord/discord_structs"
	"git-good-discord/utils"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
)

var (
	session *discordgo.Session
	details utils.DiscordDetails

	setPrefixRegex = regexp.MustCompile("!set prefix (.+)")
)

// Based on: https://github.com/bwmarrin/discordgo/blob/master/examples/pingpong/main.go
// Authorize bot for channel: https://discord.com/oauth2/authorize?client_id={APPLICATION_ID}&scope=bot
func (i Implementation) Start(errorChan chan error) {
	var err error
	details, err = utils.GetDiscordToken()
	if err != nil {
		errorChan <- err
		return
	}

	session, err = discordgo.New("Bot " + details.Token)
	if err != nil {
		errorChan <- err
		return
	}

	session.AddHandler(getMessageHandler(i))

	session.Identify.Intents = discordgo.IntentsGuildMessages

	err = session.Open()
	if err != nil {
		errorChan <- err
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	_ = session.Close()
	os.Exit(0) // Sending signal as original signal was consumed
}

func getMessageHandler(i Implementation) func (s *discordgo.Session, m *discordgo.MessageCreate) {
	return func (s *discordgo.Session, m *discordgo.MessageCreate) {
		// Ignore messages sent by bot
		if m.Author.ID == s.State.User.ID {
			return
		}

		settings, err := i.DatabaseService.GetConnection().GetChannelSettings(m.ChannelID)
		if err != nil {
			log.Println(err.Error())
			return
		}
		prefix := "!"
		if settings.Prefix != "" { prefix = strings.ToLower(settings.Prefix) }
		language := "english"
		if settings.Language != "" { language = settings.Language }

		if match := setPrefixRegex.FindStringSubmatch(strings.ToLower(m.Content)); len(match) == 2 {
			nPrefix := match[1]
			_ = i.SendMessage(discord_structs.EmbeddedMessage{Message: discord_messages.SetPrefix(i.DatabaseService, s, m, nPrefix, language)})
			return
		}

		if strings.HasPrefix(strings.ToLower(m.Content), prefix) {
			parts := strings.Split(m.Content, " ")
			command := strings.Trim(strings.ToLower(parts[0]), prefix)

			info := parts[1:]
			switch strings.ToLower(command) {
			case "command":
				err := i.SendMessage(discord_structs.EmbeddedMessage{Message: discord_structs.Message{
					ChannelID: m.ChannelID,
					Message:   fmt.Sprintf("Command: %s\nInfo: %v", command, info),
					Mentions:  []string{m.Author.Mention()},
				}})
				if err != nil {
					return
				}
			case "get":
				err := i.SendMessage(discord_structs.EmbeddedMessage{Message: discord_messages.GetGetChannel(m, prefix, language)})
				if err != nil {
					return
				}
			case "subscribe":
				err := i.SendMessage(discord_structs.EmbeddedMessage{Message: discord_messages.GetSubscribe(i.DatabaseService, i.GitlabService, m, language)})
				if err != nil { return }
			case "unsubscribe":
				err := i.SendMessage(discord_structs.EmbeddedMessage{Message: discord_messages.GetUnsubscribe(i.DatabaseService, m, language)})
				if err != nil { return }
			case "ping":
				err := i.SendMessage(discord_structs.EmbeddedMessage{Message: discord_messages.GetPing(s, m, prefix, language)})
				if err != nil {
					return
				}
			case "reload":
				err := i.SendMessage(discord_messages.GetReloadLanguage(m, language))
				if err != nil {
					return
				}
			case "language":
				err := i.SendMessage(discord_messages.GetChangeLanguage(i.DatabaseService, s, m, prefix, language))
				if err != nil {
					return
				}
			case "help":
				//Call abstact function
				//Then send result to send message
				err := i.SendMessage(discord_messages.GetHelp(s, m, prefix))
				if err != nil {
					return
				}
			default:
				err := i.SendMessage(discord_structs.EmbeddedMessage{Message: discord_structs.Message{
					ChannelID: m.ChannelID,
					Message:   fmt.Sprintf("Command: \"%s\" not recognized", command),
					Mentions:  []string{m.Author.Mention()},
				}})
				if err != nil {
					return
				}
			}
		}
	}
}