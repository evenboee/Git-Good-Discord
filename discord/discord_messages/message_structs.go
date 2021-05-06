package discord_messages

type commands struct {
	Ping ping `json:"ping"`
	GetChannel getChannel `json:"getChannel"`
}

type ping struct {
	ErrorGettingRoles string `json:"errorGettingRoles"`
	RoleNotFound string `json:"roleNotFound"`
}

type getChannel struct {
	NotSpecified string `json:"notSpecified"`
	NotRecognized string `json:"notRecognized"`
}