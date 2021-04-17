package gitlab

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// RegisterWebhook will register a Webhook for the given Gitlab Project and
// return Registration information. This function does not store the registration
// anywhere, nor does it create some sort of webhook invocation handler. That
// kind of functionality is up to the caller to implement.
func RegisterWebhook (project Project, webhook Webhook) (WebhookRegistration, error) {
	url := fmt.Sprintf("%s://%s/api/v4/projects/%d/hooks", project.Instance.Protocol, project.Instance.Host, project.ID)

	webhookJSON, err := json.Marshal(webhook)

	if err != nil {
		return WebhookRegistration{}, fmt.Errorf("could not marshal webhook json. %v", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(webhookJSON))
	req.Header.Set("content-type", "application/json")
	req.Header.Add("PRIVATE-TOKEN", project.AccessToken)

	if err != nil {
		return WebhookRegistration{}, fmt.Errorf("could not create webhook registration request. %v", err)
	}

	resp, err := client.Do(req)

	if err != nil {
		return WebhookRegistration{}, fmt.Errorf("could not POST webhook. %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		return WebhookRegistration{}, fmt.Errorf("gitlab responded with status code %s. %v", resp.Status, tryToParseErrorResponse(resp.Body))
	}

	webhookRegistration := WebhookRegistration{}
	err = json.NewDecoder(resp.Body).Decode(&webhookRegistration)
	defer resp.Body.Close()

	if err != nil {
		return WebhookRegistration{}, fmt.Errorf("could not decode webhook registration response. %v", err)
	}

	return webhookRegistration, nil
}