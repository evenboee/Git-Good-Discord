package main

import (
	"git-good-discord/http_serving"
)

func main() {
	go http_serving.StartWebHandler()
}
