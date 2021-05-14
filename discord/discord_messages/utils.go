package discord_messages

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strconv"
	"strings"
)

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
			if strings.HasPrefix(role.Name, "@") {
				return role.Name + " "
			}
			return role.Mention()
		}
	}
	return ""
}

func placeholderHandler(message string, args ...string) string {
	for i, arg := range args {
		index := strconv.Itoa(i)
		if strings.Contains(message, "{{"+index+"}}") {
			message = strings.ReplaceAll(message, "{{"+index+"}}", arg)
		}
	}
	return message
}
