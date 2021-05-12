package gitlab

import (
	"encoding/json"
	"fmt"
	"git-good-discord/gitlab/gitlab_structs"
	"git-good-discord/utils"
	"net/http"
	"net/url"
)

func (i Implementation) GetRegisteredWebhooks (project gitlab_structs.Project) ([]gitlab_structs.WebhookRegistration, error) {
	projectUrl, err := url.Parse(utils.HTTPS(project.URL))

	if err != nil {
		return nil, fmt.Errorf("could not parse project url. %v", err)
	}

	// Set correct URL parts
	projectUrl.Path = fmt.Sprintf("api/v4/projects/%d/hooks", project.ID)
	projectUrl.RawQuery = ""
	projectUrl.Fragment = ""

	url := projectUrl.String()

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("PRIVATE-TOKEN", project.AccessToken)

	if err != nil {
		return nil, fmt.Errorf("could not create GET request for webhook registrations. %v", err)
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("could not send GET request for url '%s'. %v", url, err)
	}

	var registeredWebhooks []gitlab_structs.WebhookRegistration
	err = json.NewDecoder(resp.Body).Decode(&registeredWebhooks)
	defer resp.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("could not decode response body from url '%s' to webhooks registrations. %v", url, err)
	}

	return registeredWebhooks, nil
}

func (i Implementation) DoesWebhookWithURLExist (project gitlab_structs.Project, invocationURL string) (bool, error) {
	parsedInvocationURL, err := url.Parse(utils.HTTPS(invocationURL))

	if err != nil {
		return false, fmt.Errorf("could not parse invocation URL '%s'. %v", invocationURL, err)
	}

	registeredWebhooks, err := i.GetRegisteredWebhooks(project)

	if err != nil {
		return false, fmt.Errorf("could not get registered webhooks. %v", err)
	}

	for _, webhook := range registeredWebhooks {
		webhookURL, err := url.Parse(webhook.Url)

		// Ignore error if URL is invalid
		if err != nil {
			continue
		}

		// Check if url's match
		if webhookURL.String() == parsedInvocationURL.String() {
			return true, nil
		}
	}

	return false, nil
}