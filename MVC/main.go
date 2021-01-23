package main

import (
	"WebApplications/MVC/Views"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var hoemView *Views.View
var ContactView *Views.View
var SignUpView *Views.View

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(hoemView.Render(w, nil))

}

func Contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(ContactView.Render(w, nil))
}

func SignUpPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(SignUpView.Render(w, nil))
}

func PageNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Error 404 Webpage not found</h1><p>Contact if redirectd to Anothe page </p>")
}

func main() {

	hoemView = Views.NewView("bootstrap", "Views/home.gohtml")
	ContactView = Views.NewView("bootstrap", "Views/Contact.gohtml")
	SignUpView = Views.NewView("bootstrap", "Views/SignUp.gohtml")
	r := mux.NewRouter()
	r.HandleFunc("/", Home)
	r.HandleFunc("/Contact", Contact)
	r.HandleFunc("/SignUp", SignUpPage)
	r.NotFoundHandler = http.HandlerFunc(PageNotFound)
	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
