package main

import (
	"WebApplications/MVC/Views"
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

func FAQ(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	fmt.Fprint(w, "<h1>Frequently Asked Quations</h1>")
	fmt.Fprint(w, "<ul>")
	fmt.Fprint(w, "<li>Can I GEt Free Trial Before purcahse</li>")
	fmt.Fprint(w, "<li>How do i purchase a membership for nintendo switch online </li>")
	fmt.Fprint(w, "<li>Will men=mvership plans changed automatically </li>")
	fmt.Fprint(w, "<li>Acn i change aFrom An individual Membership To a famely Membership</li>")

	fmt.Fprint(w, "</ul>")

}

func PageNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Error 404 Webpage not found</h1><p>Contact if redirectd to Anothe page </p>")
}

func main() {

	hoemView = Views.NewView("bootstrap", "Views/home.gohtml")
	ContactView = Views.NewView("bootstrap", "Views/Contact.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", Home)
	r.HandleFunc("/Contact", Contact)
	r.HandleFunc("/faq", FAQ)
	r.NotFoundHandler = http.HandlerFunc(PageNotFound)
	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
