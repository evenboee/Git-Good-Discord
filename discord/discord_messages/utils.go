package discord_messages

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func splitMessage(message string, prefix string) (string, []string) {
	parts := strings.Split(message, " ")
	command := strings.Trim(parts[0], prefix)
	info := parts[1:]
	return command, info
}

func discordMention(mentioned interface{}) string {
	switch mentioned.(type) {
	case *discordgo.User:
		user, ok := mentioned.(*discordgo.User)
		if ok {
			return user.Mention()
		}
	case *discordgo.Channel:
		channel, ok := mentioned.(*discordgo.Channel)
		if ok {
			return channel.Mention()
		}
	case *discordgo.Role:
		role, ok := mentioned.(*discordgo.Role)
		fmt.Println(role.Name)
		if ok {
			if strings.HasPrefix(role.Name, "@"){
				return role.Name + " "
			}
			return role.Mention()
		}
	}
	return ""
}
