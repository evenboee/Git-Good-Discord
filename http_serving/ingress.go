package http_serving

import (
	"net/http"
	"strings"
)

const port = "8080"

var fs = http.FileServer(http.Dir("./static"))

func StartWebHandler(errorChannel chan <- error) {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(redirectionHandler))
	err := http.ListenAndServe(":"+port, mux)
	errorChannel <- err
}

func redirectionHandler(w http.ResponseWriter, req *http.Request) {
	switch {
	case strings.Contains(strings.ToLower(req.RequestURI), "gitlab"):
		http.HandlerFunc(gitlabWebInvocationHandler).ServeHTTP(w, req)
	default:
		fs.ServeHTTP(w, req)
	}
}

func gitlabWebInvocationHandler(w http.ResponseWriter, req *http.Request) {
	//TODO: Send this to another function #9
	_,_ = w.Write([]byte("Gitlab"))
}