package gitlab_structs

// Webhook is a Webhook that can be registered with Gitlab Projects
type Webhook struct {

	// Url is the invocation URL of the webhook
	Url string `json:"url"`

	// IssuesEvents will invoke on issue events
	IssuesEvents bool `json:"issues_events"`

	// MergeRequestsEvents will invoke on merge request events
	MergeRequestsEvents bool `json:"merge_requests_events"`

	// (More event types can be added as needed)
}

// WebhookRegistration is the response given by Gitlab when a webhook has been
// successfully registered
type WebhookRegistration struct {

	// ID is the unique ID of the webhook registration for the given Gitlab Project
	ID int `json:"id"`

	// ProjectID is the ID of the project that the webhook is registered with
	ProjectID int `json:"project_id"`

	// Inherit Webhook struct
	Webhook

}

// Instance refers to a Gitlab Instance
type Instance struct {

	// Host is the hostname / IP for Gitlab Instance
	Host string `json:"host"`

	// Protocol is the protocol used when communicating with Gitlab Instance.
	// Typically, this is 'https'
	Protocol string `json:"protocol"`

}

// Project refers to a unique Gitlab Project
type Project struct {

	// Instance this Project belongs to
	Instance *Instance

	// ID is the ID given to Gitlab project
	ID int `json:"id"`

	// AccessToken required for authentication
	AccessToken string `json:"access_token"`

}

// ErrorResponse represents the error message Gitlab will send back as a http
// response
type ErrorResponse struct {

	// Message is the error message
	Message string `json:"message"`

}

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	Email string `json:"email"`
}

type ObjectAttributes struct {
	AssigneeID int `json:"assignee_id"`
	AuthorID int `json:"author_id"`
	CreatedAt string `json:"created_at"`
	Description string `json:"description"`
	MergeStatus string `json:"merge_status"`
	SourceBranch string `json:"source_branch"`
	TargetBranch string `json:"target_branch"`
	Title string `json:"title"`
}

type WebhookNotificationObject struct {
	ObjectKind string `json:"object_kind"`
}

type MergeRequestWebhookNotification struct {
	WebhookNotificationObject
	User User `json:"user"`
	ObjectAttributes ObjectAttributes `json:"object_attributes"`
}