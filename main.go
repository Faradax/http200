package main

import (
	"net/http"
	"os"
	"fmt"
	"strings"
)

var path = "/"
var useBasicAuth bool = strings.ToLower(os.Getenv("BASIC_AUTH_ENABLED")) == "true"

func allHandler(w http.ResponseWriter, r *http.Request) {
}

func checkCredentials(user string, password string) bool {
	return user == "user" && password == "password"
}

func auth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	  user, password, _ := r.BasicAuth()
	  if !checkCredentials(user, password) {
		w.Header().Set("WWW-Authenticate", "Basic realm=\"Access to the site requires auth\"")
		http.Error(w, "Unauthorized.", 401)
		return
	  }
	  handler(w, r)
	}
  }

func provideHandler(useBasicAuth bool) http.HandlerFunc {
	var handler http.HandlerFunc = allHandler
	if (useBasicAuth) {
		handler = auth(handler)
	}
	return handler
}

func main() {
	var startupLog string = "Serving http requests on port 8080"
	if (useBasicAuth) {
		startupLog = startupLog + " using basic auth"
	}
	fmt.Println(startupLog)
	var handler http.HandlerFunc = provideHandler(useBasicAuth)
	http.HandleFunc(path, handler)
	http.ListenAndServe(":8080", nil)
}