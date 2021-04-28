package http_serving

import (
	"git-good-discord/gitlab/gitlab_interfaces"
	"log"
	"net/http"
	"strings"
)

const port = "8080"

var (
	Gitlab gitlab_interfaces.Interface

	fs = http.FileServer(http.Dir("./static"))
)

func StartWebHandler(errorChannel chan <- error) {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(redirectionHandler))
	err := http.ListenAndServe(":"+port, mux)
	errorChannel <- err
}

func redirectionHandler(w http.ResponseWriter, req *http.Request) {
	switch {
	case strings.Contains(strings.ToLower(req.RequestURI), "gitlab"):
		// Wrap handler with error handling and handle request
		http.HandlerFunc(
			func(writer http.ResponseWriter, request *http.Request) {
				err := Gitlab.HandleWebhookNotificationHTTP(w, request)

				if err != nil {
					// Don't want to crash the program by sending error to error channel, so the
					// error gets logged instead
					log.Printf("error handling webhook notification. %v\n", err)
				}
			},
		).ServeHTTP(w, req)
	default:
		fs.ServeHTTP(w, req)
	}
}