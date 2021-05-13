package http_serving

import (
	"git-good-discord/gitlab/gitlab_interfaces"
	"log"
	"net/http"
	"strings"
)

const port = "8080"

type Implementation struct {
	GitlabService gitlab_interfaces.Interface
}

var fs = http.FileServer(http.Dir("./static"))

func (i Implementation) Start(errorChannel chan <- error) {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(getRedirectionHandler(i)))

	log.Println("HTTP Web Handler started")
	err := http.ListenAndServe(":"+port, mux)
	errorChannel <- err
}

func getRedirectionHandler(i Implementation) func (http.ResponseWriter, *http.Request) {
	return func (w http.ResponseWriter, req *http.Request) {
		switch {
		case strings.Contains(strings.ToLower(req.RequestURI), "gitlab"):
			http.HandlerFunc(getGitlabWebhookNotificationHandler(i)).ServeHTTP(w, req)
		default:
			fs.ServeHTTP(w, req)
		}
	}
}

func getGitlabWebhookNotificationHandler (i Implementation) func(http.ResponseWriter, *http.Request) {
	// Wrap handler with error handling and handle request
	return func(w http.ResponseWriter, req *http.Request) {
		err := i.GitlabService.HandleWebhookNotificationHTTP(w, req)

		if err != nil {
			// Don't want to crash the program by sending error to error channel, so the
			// error gets logged instead
			log.Printf("error handling webhook notification. %v\n", err)
		}
	}
}