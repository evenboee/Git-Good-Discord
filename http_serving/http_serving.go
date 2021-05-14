package http_serving

import "git-good-discord/gitlab/gitlab_interfaces"

// Implementation for the http_serving_interfaces.WebHandler
type Implementation struct {

	// GitlabService required by Implementation
	GitlabService gitlab_interfaces.Interface
}
