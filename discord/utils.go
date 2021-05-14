package discord

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

func memberIsAdmin(m *discordgo.MessageCreate, s *discordgo.Session,) bool {
	roles, err := s.GuildRoles(m.GuildID)
	if err != nil {
		log.Printf("Bot.go: %v\n", err)
		return false
	}

	adminID := ""
	for _, role := range roles {
		if role.Name == "Admin" { adminID = role.ID }
	}
	if adminID == "" { return false }

	for _, role := range m.Member.Roles {
		if role == adminID { return true }
	}
	return false
}

func splitMessage(message string, prefix string) (string, []string) {
	parts := strings.Split(message, " ")
	command := strings.Trim(parts[0], prefix)
	info := parts[1:]
	return command, info
}