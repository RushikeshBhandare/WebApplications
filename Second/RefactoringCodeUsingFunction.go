package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcoem to my page</h1>")
	} else if r.URL.Path == "/Contact" {
		fmt.Fprint(w, "To get in touch please send an email to <a href=\"google.com\" >Gmail</a>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>We could not found your page for :(</h1> <p>plese contact us if any queries</p>")

	}

}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome back</h1>")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "</h1>To get in touch please send an email to <a href=\"google.com\" >Gmail</a>")
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
	r := mux.NewRouter()
	r.HandleFunc("/", Home)
	r.HandleFunc("/Contact", Contact)
	r.HandleFunc("/faq", FAQ)
	r.NotFoundHandler = http.HandlerFunc(PageNotFound)
	http.ListenAndServe(":3000", r)
}
