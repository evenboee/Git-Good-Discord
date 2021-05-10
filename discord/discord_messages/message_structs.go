package discord_messages

type commands struct {
	Ping           Ping           `json:"ping"`
	GetChannel     GetChannel     `json:"getChannel"`
	ReloadLanguage ReloadLang     `json:"reloadLang"`
	ChangeLanguage ChangeLanguage `json:"changeLang"`
	Language       string         `json:"language"`
	SetPrefix      setPrefix      `json:"setPrefix"`
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

type setPrefix struct {
	NotAuthorized string `json:"notAuthorized"`
	Successful    string `json:"successful"`
}
