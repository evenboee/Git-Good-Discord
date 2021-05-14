package discord

import (
	"git-good-discord/database/database_interfaces"
	"git-good-discord/discord/discord_messages"
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

func getMessageHandler(i Implementation) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// Ignore messages sent by bot
		if m.Author.ID == s.State.User.ID {
			return
		}
		prefix, language := getSettings(i.DatabaseService.GetConnection(), m.ChannelID)
		if setPrefix(i, s, m, language) {
			return
		}

		if strings.HasPrefix(strings.ToLower(m.Content), prefix) {
			commandHandler(i, s, m, prefix, language)
		}
	}
}

func getSettings(conn database_interfaces.DatabaseConnection, channel_id string) (string, string) {
	prefix := "!"
	language := "english"

	settings, err := conn.GetChannelSettings(channel_id)
	if err != nil {
		log.Println(err.Error())
	} else {
		if settings.Prefix != "" {
			prefix = strings.ToLower(settings.Prefix)
		}
		if settings.Language != "" {
			language = settings.Language
		}
	}
	return prefix, language
}

func setPrefix(i Implementation, s *discordgo.Session, m *discordgo.MessageCreate, language string) bool {
	if match := regexp.MustCompile("!set prefix (.+)").FindStringSubmatch(strings.ToLower(m.Content)); len(match) == 2 {
		if memberIsAdmin(m, s) {
			newPrefix := match[1]
			_ = i.SendMessage(setPrefixHandler(newPrefix, language, i.DatabaseService, m))
		} else {
			_ = i.SendMessage(discord_messages.NotAuthorizedMessage(language, "SetPrefix", m))
		}
		return true
	}
	return false
}
