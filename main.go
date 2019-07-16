package main

import (
	"fmt"
	"myGo/Go-web/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Page not found</h1>")
}
func main() {
	staticC := controllers.NewStatic()
	userC := controllers.NewUsers()

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.Handle("/signup", userC.NewView).Methods("GET")
	// r.HandleFunc("/signup", userC.New).Methods("GET")
	r.HandleFunc("/signup", userC.Create).Methods("POST")
	http.ListenAndServe(":3000", r)
}

// This function is just for development mode. We will handle different later this error in production mode.
func must(err error) {
	if err != nil {
		panic(err)
	}
}
