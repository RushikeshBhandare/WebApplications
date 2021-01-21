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
	fmt.Fprint(w, " <h1>Error 404 </h1>To get in touch please send an email to <a href=\"google.com\" >Gmail</a>")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Home)
	r.HandleFunc("/Contact", Contact)
	http.ListenAndServe(":3000", r)
}
