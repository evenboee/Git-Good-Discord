package discord_messages

type commands struct {
	Ping                     Ping                     `json:"ping"`
	GetChannel               GetChannel               `json:"getChannel"`
	ReloadLanguage           ReloadLang               `json:"reloadLang"`
	ChangeLanguage           ChangeLanguage           `json:"changeLang"`
	HelpCommand              HelpCommand              `json:"helpCommand"`
	SetLanguagePrefix        SetLanguagePrefix        `json:"setPrefix"`
	Language                 string                   `json:"language"`
	NotificationMergeRequest NotificationMergeRequest `json:"notifyMR"`
	Subscribe                Subscribe                `json:"subscribe"`
	Unsubscribe              Unsubscribe              `json:"unSubscribe"`
}

type Ping struct {
	ErrorGettingRoles string `json:"errorGettingRoles"`
	RoleNotFound      string `json:"roleNotFound"`
}

type GetChannel struct {
	NotSpecified  string `json:"notSpecified"`
	NotRecognized string `json:"notRecognized"`
}

type ReloadLang struct {
	ErrorReloading       string `json:"errorReloading"`
	SuccessfullyReloaded string `json:"successfullyReloaded"`
}

type ChangeLanguage struct {
	NoParam         string `json:"noParam"`
	InvalidLanguage string `json:"invalidLang"`
	Successful      string `json:"successful"`
	DatabaseSetFail string `json:"databaseSetFail"`
	NotAuthorized   string `json:"notAuthorized"`
}

type SetLanguagePrefix struct {
	NotAuthorized string `json:"notAuthorized"`
	Successful    string `json:"successful"`
}

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

type Unsubscribe struct {
	DatabaseRemoveFail string `json:"databaseRemoveFail"`
	PathFormatError    string `json:"pathFormatError"`
	PartsError         string `json:"partsError"`
	Successful         string `json:"successful"`
}

type NotificationMergeRequest struct {
	Success string `json:"success"`
}
