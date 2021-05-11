package discord_messages

type commands struct {
	Ping              Ping              `json:"ping"`
	GetChannel        GetChannel        `json:"getChannel"`
	ReloadLanguage    ReloadLang        `json:"reloadLang"`
	ChangeLanguage    ChangeLanguage    `json:"changeLang"`
	HelpCommand       HelpCommand       `json:"helpCommand"`
	SetLanguagePrefix SetLanguagePrefix `json:"setPrefix"`
	Language          string            `json:"language"`
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
	Get       string `json:"get"`
	Ping      string `json:"ping"`
	Reload    string `json:"reload"`
	Language  string `json:"language"`
	Help      string `json:"help"`
	AdminOnly string `json:"adminOnly"`
	SetPrefix string `json:"setPrefix"`
}
