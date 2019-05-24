package main

import (
	"fmt"
	"net/http"
)

// This function is call anytime, anyone comes to the webserver(website)
// w http.ResponseWriter, this is where we write the response to our user
// r *http.Request has information about the request the user is making
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	// the first argument, we tell where we want to print to
	fmt.Fprint(w, "<h1>Welcome to my awesome site</h1>")
}

func main() {
	// the first argument is the path we want to match, and everytime is
	// match (a request), the handleFunc is called
	// "/" means any path, so no matter which path is request in the website
	// the handlerFunc is going to be called
	http.HandleFunc("/", handlerFunc)
	// here is where we start the server
	// the nil in this case means that we want to use the build in server
	// we are listening in the port 3000
	http.ListenAndServe(":3000", nil)
}
