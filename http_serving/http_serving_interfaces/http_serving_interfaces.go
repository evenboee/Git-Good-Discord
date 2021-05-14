package http_serving_interfaces

// WebHandler is an interface for the web handler
type WebHandler interface {

	// Start starts the web handler and sends error through errorChannel if an error
	// occurs
	Start(errorChannel chan<- error)
}
