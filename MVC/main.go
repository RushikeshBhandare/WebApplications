package main

import (
	"WebApplications/MVC/Views"
	"WebApplications/MVC/controlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var hoemView *Views.View
var ContactView *Views.View

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(hoemView.Render(w, nil))

}

func Contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(ContactView.Render(w, nil))
}

func PageNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Error 404 Webpage not found</h1><p>Contact if redirectd to Anothe page </p>")
}

func main() {

	hoemView = Views.NewView("bootstrap", "Views/home.gohtml")
	ContactView = Views.NewView("bootstrap", "Views/Contact.gohtml")
	UserC := controlers.NewUser()

	r := mux.NewRouter()
	r.HandleFunc("/", Home).Methods("GET")
	r.HandleFunc("/Contact", Contact).Methods("GET")
	r.HandleFunc("/SignUp", UserC.New).Methods("GET")
	r.HandleFunc("/SignUp", UserC.Create).Methods("POST")

	r.NotFoundHandler = http.HandlerFunc(PageNotFound)
	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
