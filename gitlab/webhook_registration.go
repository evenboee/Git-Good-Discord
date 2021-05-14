package gitlab

import (
	"bytes"
	"encoding/json"
	"fmt"
	"git-good-discord/gitlab/gitlab_structs"
	"git-good-discord/utils"
	"net/http"
	"net/url"
)

func (i Implementation) RegisterWebhook(project gitlab_structs.Project, webhook gitlab_structs.Webhook) (gitlab_structs.WebhookRegistration, error) {
	//If webhook exists or there is an error
	ok, err := i.DoesWebhookWithURLExist(project, webhook.Url)
	if err != nil {
		return gitlab_structs.WebhookRegistration{}, err
	} else if ok {
		return gitlab_structs.WebhookRegistration{}, fmt.Errorf("webhook is already registered. %v", err)
	}

	projectUrl, err := url.Parse(utils.HTTPS(project.URL))

	if err != nil {
		return gitlab_structs.WebhookRegistration{}, fmt.Errorf("could not parse project url. %v", err)
	}

	// Set correct URL parts
	projectUrl.Path = fmt.Sprintf("api/v4/projects/%d/hooks", project.ID)
	projectUrl.RawQuery = ""
	projectUrl.Fragment = ""

	url := projectUrl.String()

	webhookJSON, err := json.Marshal(webhook)

	if err != nil {
		return gitlab_structs.WebhookRegistration{}, fmt.Errorf("could not marshal webhook json. %v", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(webhookJSON))
	req.Header.Set("content-type", "application/json")
	req.Header.Add("PRIVATE-TOKEN", project.AccessToken)

	if err != nil {
		return gitlab_structs.WebhookRegistration{}, fmt.Errorf("could not create webhook registration request. %v", err)
	}

	resp, err := client.Do(req)

	if err != nil {
		return gitlab_structs.WebhookRegistration{}, fmt.Errorf("could not POST webhook. %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		return gitlab_structs.WebhookRegistration{}, fmt.Errorf("gitlab responded with status code %s. %v", resp.Status, tryToParseErrorResponse(resp.Body))
	}

	webhookRegistration := gitlab_structs.WebhookRegistration{}
	err = json.NewDecoder(resp.Body).Decode(&webhookRegistration)
	defer resp.Body.Close()

	if err != nil {
		return gitlab_structs.WebhookRegistration{}, fmt.Errorf("could not decode webhook registration response. %v", err)
	}

	return webhookRegistration, nil
}

func (i Implementation) GetWebhookInvocationURL(discordChannelID string) (string, error) {
	ip, err := utils.GetIP()
	if err != nil {
		return "", err
	}
	ip = utils.HTTPS(ip, true)
	parsedURL, err := url.Parse(ip)

	if err != nil {
		return "", fmt.Errorf("could not parse url '%s'. %v", ip, err)
	}

	parsedURL.Path = fmt.Sprintf("/gitlab/%s", discordChannelID)
	parsedURL.RawQuery = ""
	parsedURL.Fragment = ""

	return parsedURL.String(), nil
}
