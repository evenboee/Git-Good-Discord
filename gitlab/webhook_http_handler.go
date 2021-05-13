package gitlab

import (
	"encoding/json"
	"fmt"
	"git-good-discord/gitlab/gitlab_interfaces"
	"git-good-discord/gitlab/gitlab_structs"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (i Implementation) HandleWebhookNotificationHTTP(w http.ResponseWriter, req *http.Request) error {
	pathSplit := strings.Split(req.URL.Path, "/")

	// 3 is minimum split length
	if len(pathSplit) < 3 {
		return fmt.Errorf("not enough path separators (/) to get discord channel id from url path '%s'", req.URL.Path)
	}

	// Assumes that index 2 is discord Channel ID
	// Follows format '/gitlab/{:discord_channel_id}'
	discordChannelID := pathSplit[2]

	// Get secret token from header
	secretToken := req.Header.Get("X-Gitlab-Token")

	notification := gitlab_structs.WebhookNotification{}
	err := json.NewDecoder(req.Body).Decode(&notification)
	defer req.Body.Close()

	if err != nil {
		return fmt.Errorf("could not unmarshal webhook notification body. %v", err)
	}

	err = checkSecurityToken(i, secretToken, discordChannelID, notification.Project.URL, strconv.Itoa(notification.Project.ID))

	if err != nil {
		return err
	}

	// Check if unsupported notification type
	switch notification.ObjectKind {
	case gitlab_interfaces.NotificationMergeRequest:
		break
	case gitlab_interfaces.NotificationIssue:
		break
	default:
		return fmt.Errorf("received unsupported webhook notification type '%s'", notification.ObjectKind)
	}

	// Pass notification to abstraction layer
	i.AbstractionService.HandleGitlabNotification(notification, discordChannelID)

	return nil
}

func checkSecurityToken(i Implementation, token string, discord_channel_id string, url string, repo_id string) error {
	gitlab_instance := strings.Split(url, "/")[2]

	actualToken, err := i.DatabaseService.GetConnection().GetSecurityToken(discord_channel_id, gitlab_instance, repo_id)

	if err != nil {
		return err
	}
	if actualToken == "" {
		log.Printf("No security token registered for this request. Channel ID: %v, URL: %v, REPO id: %v\n", discord_channel_id, url, repo_id)
		return nil
	}

	if token != actualToken {
		return fmt.Errorf("wrong secret token provided in request: '%s'", token)
	}
	return nil
}
