package discord_messages

// commands contains the bot commands
type commands struct {
	Ping              			Ping              			`json:"ping"`
	GetChannel        			GetChannel        			`json:"getChannel"`
	ReloadLanguage    			ReloadLang        			`json:"reloadLang"`
	ChangeLanguage    			ChangeLanguage    			`json:"changeLang"`
	HelpCommand       			HelpCommand       			`json:"helpCommand"`
	SetLanguagePrefix 			SetLanguagePrefix 			`json:"setPrefix"`
	Language          			string            			`json:"language"`
	NotificationMergeRequest 	NotificationMergeRequest 	`json:"notifyMR"`
	NotificationIssue 			NotificationIssue 			`json:"notifyIssue"`
	Subscribe         			Subscribe         			`json:"subscribe"`
	Unsubscribe       			Unsubscribe       			`json:"unSubscribe"`
	Errors						Errors						`json:"errors"`
	SetAccessToken              SetAccessToken              `json:"setAccessToken"`
	Subscriptions               Subscriptions               `json:"subscriptions"`
}

// Ping contains Ping-related words
type Ping struct {
	ErrorGettingRoles string `json:"errorGettingRoles"`
	RoleNotFound      string `json:"roleNotFound"`
}

// GetChannel contains GetChannel-related words
type GetChannel struct {
	NotSpecified  string `json:"notSpecified"`
	NotRecognized string `json:"notRecognized"`
}

// ReloadLang contains ReloadLang-related words
type ReloadLang struct {
	ErrorReloading       string `json:"errorReloading"`
	SuccessfullyReloaded string `json:"successfullyReloaded"`
}

// ChangeLanguage contains ChangeLanguage-related words
type ChangeLanguage struct {
	NoParam         string `json:"noParam"`
	InvalidLanguage string `json:"invalidLang"`
	Successful      string `json:"successful"`
	DatabaseSetFail string `json:"databaseSetFail"`
	NotAuthorized   string `json:"notAuthorized"`
}

// SetLanguagePrefix contains SetLanguagePrefix-related words
type SetLanguagePrefix struct {
	NotAuthorized string `json:"notAuthorized"`
	Successful    string `json:"successful"`
}

// HelpCommand contains HelpCommand-related words
type HelpCommand struct {
	Get         string `json:"get"`
	Ping        string `json:"ping"`
	Reload      string `json:"reload"`
	Language    string `json:"language"`
	Help        string `json:"help"`
	AdminOnly   string `json:"adminOnly"`
	SetPrefix   string `json:"setPrefix"`
	Subscribe   string `json:"subscribe"`
	Unsubscribe string `json:"unsubscribe"`
}

// Subscribe contains Subscribe-related words
type Subscribe struct {
	DatabaseAddFail              string `json:"databaseAddFail"`
	Successful                   string `json:"successful"`
	PathFormatError              string `json:"pathFormatError"`
	TokenGenerationFail          string `json:"tokenGenerationFail"`
	InvocationURLFail            string `json:"invocationURLFail"`
	RepoIDFormatError            string `json:"repoIDFormatError"`
	WebhookRegistrationError     string `json:"webhookRegistrationError"`
	DatabaseAddSecurityTokenFail string `json:"databaseAddSecurityTokenFail"`
	AccessTokenFail              string `json:"accessTokenFail"`
}

// Unsubscribe contains Unsubscribe-related words
type Unsubscribe struct {
	DatabaseRemoveFail string `json:"databaseRemoveFail"`
	PathFormatError    string `json:"pathFormatError"`
	PartsError         string `json:"partsError"`
	Successful         string `json:"successful"`
}

// NotificationMergeRequest contains NotificationMergeRequest-related words
type NotificationMergeRequest struct {
	Success string `json:"success"`
}

// NotificationIssue contains NotificationIssue-related words
type NotificationIssue struct {
	Success string `json:"success"`
}

// Errors contains error-related words
type Errors struct {
	CommandNotRecognized string `json:"commandNotRecognized"`
}

// SetAccessToken contains messages related to setting the access token of a gitlab project
type SetAccessToken struct {
	NotAuthorized    string `json:"notAuthorized"`
	WrongParts       string `json:"wrongParts"`
	WrongPath        string `json:"wrongPath"`
	PathElementEmpty string `json:"pathElementEmpty"`
	AddTokenFail     string `json:"addTokenFail"`
	Successful       string `json:"successful"`
}

type Subscriptions struct {
	DatabaseFail string `json:"databaseFail"`
	Successful   string `json:"successful"`
}