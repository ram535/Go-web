package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awesome site</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send a email to <a href=\"mailto:support@foo.com\">suppurt@foo.com</a>")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Page not found</h1>")
}
func main() {
	r := mux.NewRouter()
	// this is use when no route match
	// it is a build in variable of type http.Handler
	// it use by default a predifine 404 page
	// we change it to use a custom 404 page
	//
	// type HandlerFunc func(ResponseWriter, *Request)
	// func (f HandlerFunc) ServerHTTP(w ResponseWriter, r *Rquest)
	// type Handler interface {
	//    ServerHTTP(ResponseWriter, *Request)
	// }
	//
	// NotFoundHandler has to be of type http.Handler
	// that's why we transform it with http.HandlerFunc()
	// I understand we can transform it because the underline
	// type of HandlerFunc is the same as the signature of notFound
	// so by transforming notFound to http.HandlerFunc, it also become
	// of type http.Handler because http.HandlerFunc has the ServerHTTP()
	// to be able to also be of the type http.Handler because it match
	// the interface of http.Handler
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)
}
