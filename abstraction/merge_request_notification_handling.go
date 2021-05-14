package abstraction

import (
	"git-good-discord/database/database_structs"
	"git-good-discord/discord/discord_messages"
	"git-good-discord/gitlab/gitlab_structs"
	"git-good-discord/utils"
	"log"
	"net/url"
	"strconv"
	"strings"
)

func (i Implementation) HandleGitlabNotification(notification gitlab_structs.WebhookNotification, discordChannelID string) {
	parsedURL, err := url.Parse(utils.HTTPS(notification.Project.URL))

	if err != nil {
		log.Printf("could not parse project url when handling merge request notification. %v", err)
		return
	}

	// Get hostname of gitlab instance
	gitlabInstance := parsedURL.Hostname()

	repoID := notification.Project.ID

	// Unique usernames in form of "Index -> Username" map Index can be prefixed by
	// "id:" or "name:" The advantage of this in comparison to normal []string is
	// that it requires fewer "ID -> Username" lookups
	uniqueUsernames := make(map[string]string, 3)

	// Helper functions to add unique usernames with the fewest amount of "ID -> Username" lookups
	addUsernamesIfAbsent(&uniqueUsernames, i, parsedURL.String(), notification.ObjectAttributes.AssigneeID, notification.ObjectAttributes.AuthorID)
	addUsernameIfAbsent(&uniqueUsernames, notification.User.Username)

	// Get all interested subscribers for the given usernames
	interestedSubscribers := getInterestedSubscribers(&uniqueUsernames, i, gitlabInstance, discordChannelID, strconv.Itoa(repoID))

	// Send message to notify subscribers
	i.DiscordService.SendMessage(discord_messages.NotifySubscribers(discordChannelID, interestedSubscribers, notification))
}

func getInterestedSubscribers (uniqueUsernames *map[string]string, i Implementation, gitlabInstance, discordChannelID string, repoID string) []database_structs.Subscriber {
	// Interested subscribers map in form of "Discord ID -> Subscriber"
	interestedSubscribersMap := make(map[string]database_structs.Subscriber)

	databaseConnection := i.DatabaseService.GetConnection()

	// Find all subscribers that have subscribed (are interested) in any of the Gitlab usernames
	for _, uniqueUsername := range *uniqueUsernames {
		subscribers, err := databaseConnection.GetSubscribers(discordChannelID, gitlabInstance, repoID, uniqueUsername)

		if err != nil {
			// Log error, but continue
			log.Printf("could not ping subscribers when handling merge request notification. %v", err)
			continue
		}

		// Check if any discord users are interested in notifications for the current gitlab username
		for _, sub := range subscribers {
			// Check if subscriber is actually interested in merge requests
			if sub.MergeRequests {
				// Check if subscriber already has been registered as interested
				if _, exists := interestedSubscribersMap[sub.DiscordUserId]; !exists {
					// Add interested subscriber
					interestedSubscribersMap[sub.DiscordUserId] = sub
				}
			}
		}
	}

	// Convert interested subscribers map to slice
	interestedSubscribers := make([]database_structs.Subscriber, 0, len(interestedSubscribersMap))

	for _, sub := range interestedSubscribersMap {
		interestedSubscribers = append(interestedSubscribers, sub)
	}

	return interestedSubscribers
}

func addUsernameIfAbsent (uniqueUsernames *map[string]string, username string) {
	// Return if username exists in map
	for _, mapUsername := range *uniqueUsernames {
		if strings.EqualFold(username, mapUsername) {
			return
		}
	}

	(*uniqueUsernames)["name:" + username] = username
}

func addUsernamesIfAbsent (uniqueUsernames *map[string]string, i Implementation, url string, ids ...int) {
	// Only retain unique ids
	ids = utils.ConvertToUniqueIntSlice(ids)

	// Add all unique usernames from ids
	// Only fetch user data when necessary
	for _, id := range ids {
		// Ignore if id is 0
		if id == 0 {
			continue
		}

		index := "id:" + strconv.Itoa(id)

		if _, exists := (*uniqueUsernames)[index]; !exists {
			user, err := i.GitlabService.GetUserByID(url, id)

			if err != nil {
				// Log error, but continue
				log.Printf("could not get gitlab user info using id '%d'. %v", id, err)
				continue
			}

			(*uniqueUsernames)[index] = user.Username
		}
	}
}