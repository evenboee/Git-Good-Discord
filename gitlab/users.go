package gitlab

import (
	"encoding/json"
	"fmt"
	"git-good-discord/gitlab/gitlab_structs"
	"net/http"
	"net/url"
	"strings"
)

func (i Implementation) GetUserByUsername (instanceURL string, username string) (gitlab_structs.User, error) {
	parsedURL, err := url.Parse(instanceURL)

	if err != nil {
		return gitlab_structs.User{}, fmt.Errorf("could not parse project url. %v", err)
	}

	// Set correct URL parts
	parsedURL.Path = "api/v4/users"
	parsedURL.RawQuery = fmt.Sprintf("username=%s", username)
	parsedURL.Fragment = ""

	getURL := parsedURL.String()

	resp, err := http.Get(getURL)

	if err != nil {
		return gitlab_structs.User{}, fmt.Errorf("could not do GET request for getting gitlab user info. %v", err)
	}

	var users []gitlab_structs.User

	err = json.NewDecoder(resp.Body).Decode(&users)
	defer resp.Body.Close()

	if err != nil {
		return gitlab_structs.User{}, fmt.Errorf("could not decode gitlab user JSON for url '%s'. %v", getURL, err)
	}

	// Return user if username is a case-insensitive match
	for _, user := range users {
		if strings.EqualFold(user.Username, username) {
			return user, nil
		}
	}

	// Return error if username was not found
	return gitlab_structs.User{}, fmt.Errorf("could not find user with username '%s'", username)
}

func (i Implementation) GetUserByID(instanceURL string, id int) (gitlab_structs.User, error) {
	parsedURL, err := url.Parse(instanceURL)

	if err != nil {
		return gitlab_structs.User{}, fmt.Errorf("could not parse project url. %v", err)
	}

	// Set correct URL parts
	parsedURL.Path = fmt.Sprintf("api/v4/users/%d", id)
	parsedURL.RawQuery = ""
	parsedURL.Fragment = ""

	getURL := parsedURL.String()

	resp, err := http.Get(getURL)

	if err != nil {
		return gitlab_structs.User{}, fmt.Errorf("could not do GET request for getting gitlab user info. %v", err)
	}

	user := gitlab_structs.User{}
	err = json.NewDecoder(resp.Body).Decode(&user)
	defer resp.Body.Close()

	if err != nil {
		return gitlab_structs.User{}, fmt.Errorf("could not decode gitlab user JSON for url '%s'. %v", getURL, err)
	}

	return user, nil
}