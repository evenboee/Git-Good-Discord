package discord

import (
	"fmt"
	"git-good-discord/discord/discord_messages"
	"git-good-discord/discord/discord_structs"
	"git-good-discord/utils"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
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

	session.AddHandler(messageHandler)

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

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages sent by bot
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "!") {
		parts := strings.Split(m.Content, " ")
		command := strings.Trim(parts[0], "!")
		info := parts[1:]
		switch command {
		case "command":
			err := GetImplementation().SendMessage(discord_structs.EmbeddedMessage{Message: discord_structs.Message{
				ChannelID: m.ChannelID,
				Message:   fmt.Sprintf("Command: %s\nInfo: %v", command, info),
				Mentions:  []string{m.Author.Mention()},
			}})
			if err != nil {
				return
			}
		case "get":
			err := GetImplementation().SendMessage(discord_structs.EmbeddedMessage{Message: discord_messages.GetChannel(m, "!")})
			if err != nil {
				return
			}
		case "ping":
			err := GetImplementation().SendMessage(discord_structs.EmbeddedMessage{Message: discord_messages.Ping(s, m, "!")})
			if err != nil {
				return
			}
		default:
			err := GetImplementation().SendMessage(discord_structs.EmbeddedMessage{Message: discord_structs.Message{
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
