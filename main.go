package main

import (
	"fmt"
	"myGo/Go-web/views"
	"net/http"

	"github.com/gorilla/mux"
)

// we could have put this variable inside the func home
// but that would create a home page everytime the
// func home is called
var homeView *views.View
var contactView *views.View

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// homeView.Layout was define as "bootstrap" down there in main(). So this
	// method look for the "bootstrap template". I define it in bootstrap.gohtml
	// with {{define bootstrap}}
	err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)
	if err != nil {
		panic(err)
	}

}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Page not found</h1>")
}
func main() {
	// we are passing the home.gohtml file to the views.NewView() method
	// So this method will parse this file with the other files (booktabs.gohtml, footer.gohtml, navbar.gohtml)
	// So the program is going to be aware which files exist
	// So like noew we can call the "bootsrap template"
	// above in the homeView.Template.Executetemplate() method
	// and this one can call the "home template, navbar template and footer template"
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)
}
