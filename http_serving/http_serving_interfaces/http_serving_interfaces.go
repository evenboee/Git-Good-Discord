package http_serving_interfaces

type WebHandler interface {
	Start(errorChannel chan <- error)
}