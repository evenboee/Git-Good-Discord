package discord_messages

type commands struct {
	Ping ping `json:"ping"`
	GetChannel getChannel `json:"getChannel"`
	ReloadLanguage reloadLang `json:"reloadLang"`
	ChangeLanguage changeLang `json:"changeLang"`
	Language string `json:"language"`
}

type ping struct {
	ErrorGettingRoles string `json:"errorGettingRoles"`
	RoleNotFound string `json:"roleNotFound"`
}

type getChannel struct {
	NotSpecified string `json:"notSpecified"`
	NotRecognized string `json:"notRecognized"`
}

type reloadLang struct {
	ErrorReloading string `json:"errorReloading"`
	SuccessfullyReloaded string `json:"successfullyReloaded"`
}

type changeLang struct {
	NoParam string `json:"no_param"`
	InvalidLanguage string `json:"invalidLang"`
	Successful string `json:"successful"`
}