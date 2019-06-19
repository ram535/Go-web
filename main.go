package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my awesome site</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch, please send a email to <a href=\"mailto:support@foo.com\">suppurt@foo.com</a>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>Could not find page</h1>")
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", r)
}
