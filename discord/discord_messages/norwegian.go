package discord_messages

func (n Norwegian) getChannel() []string {
	return []string{
		"Spesifisier hva du ønsker å få!",
		"Klarer ikke å gjenkjenne: ",
	}
}

func (n Norwegian) ping() []string {
	return []string{
		"Error: Kunne ikke få tak i rollene",
		"Vi fant ikke rollen",
	}
}
