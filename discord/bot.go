package discord

import (
	"errors"
	"fmt"
	"git-good-discord/discord/discord_structs"
	"git-good-discord/utils"
	"github.com/bwmarrin/discordgo"
	"log"
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
		errorChan<-err
		return
	}

	session, err = discordgo.New("Bot " + details.Token)
	if err != nil {
		errorChan<-err
		return
	}

	session.AddHandler(messageHandler)

	session.Identify.Intents = discordgo.IntentsGuildMessages

	err = session.Open()
	if err != nil {
		errorChan<-err
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	_ = session.Close()
	os.Exit(0) // Sending signal as original signal was consumed
}


func (i Implementation) SendMessage(msg discord_structs.Message) error {
	if session == nil { return errors.New("session is not open") }
	_, err := session.ChannelMessageSend(msg.ChannelID, msg.Content)
	return err
}


func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages sent by bot
	if m.Author.ID == s.State.User.ID { return }

	if strings.HasPrefix(m.Content, "!") {
		parts := strings.Split(m.Content, " ")
		command := strings.Trim(parts[0], "!")
		info := parts[1:]
		msg := ""
		switch command {
		case "command":
			msg = fmt.Sprintf("Command: %s\nInfo: %v", command, info)
			break
		case "get":
			if len(info) == 0 {
				msg = "Specify what you want to get"
			} else {
				if info[0] == "channel_id" {
					msg = "Channel_id: " + m.ChannelID
				} else {
					msg = info[0] + " not recognized"
				}
			}
		case "ping":
			rl := ""
			if len(info) == 0 {
				rl = "everyone"
			} else { rl = info[0] }

			roles, err := s.GuildRoles(m.GuildID)
			if err != nil {
				msg = "Error: Failed to get roles"
			}
			for _, role := range roles {
				if strings.Contains(strings.ToLower(role.Name), rl) {
					msg = role.Mention()
					break
				}
			}
			if msg == "" {
				msg = "Role \"" + rl + "\" not found"
			}
			break
		default:
			msg = fmt.Sprintf("Command: \"%s\" not recognized", command)
		}

		_, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s\n%s", m.Author.Mention(), msg))
		if err != nil {
			log.Fatalln("Failed to send message")
		}
	}
}
