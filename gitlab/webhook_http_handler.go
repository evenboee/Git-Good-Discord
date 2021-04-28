package gitlab

import (
	"encoding/json"
	"fmt"
	"git-good-discord/gitlab/gitlab_structs"
	"io/ioutil"
	"net/http"
)

// NotificationMergeRequest Tried to use the same syntax as http.StatusX constants
const NotificationMergeRequest = "merge_request"

func (i Implementation) HandleWebhookNotificationHTTP(w http.ResponseWriter, req *http.Request) error {
	// Originally, I thought using json.NewDecoder(req.Body) would be cleaner
	// Unfortunately, due to the way Gitlab notifications are structured (different
	// types have different structures),
	// I couldn't find a better way than to unmarshal twice:
	// 1. Unmarshal to determine what type of notification
	// 2. Unmarshal to the actual notification type
	body, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	if err != nil {
		return fmt.Errorf("could not read request body. %v", err)
	}

	notificationObject := gitlab_structs.WebhookNotificationObject{}
	err = json.Unmarshal(body, &notificationObject)

	if err != nil {
		return fmt.Errorf("could not unmarshal webhook notification body. %v", err)
	}

	// Unmarshal to correct notification object and pass it to Abstraction layer if
	// supported. Otherwise, return an error.
	switch notificationObject.ObjectKind {
	case NotificationMergeRequest:
		notification := gitlab_structs.MergeRequestWebhookNotification{}
		err = json.Unmarshal(body, &notification)

		if err != nil {
			return fmt.Errorf("could not unmarshal webhook notification body as merge request notification. %v", err)
		}

		Abstraction.HandleGitlabMergeRequestNotification(notification)
		break

	default:
		return fmt.Errorf("received unsupported webhook notification type '%s'", notificationObject.ObjectKind)
	}

	return nil
}