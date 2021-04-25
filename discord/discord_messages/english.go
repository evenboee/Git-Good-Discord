package discord_messages

func (e English) getChannel() []string {
	return []string {
		"Specify what you want to get",
		"Could not recognize: ",
	}
}

func (e English) ping() []string {
	return []string{
		"Error: Failed to get roles",
		"Could not find role ",
	}
}
