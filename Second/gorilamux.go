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

func main() {

	m := mux.NewRouter()
	m.HandleFunc("/", handleFunc)
	http.ListenAndServe(":3000", m)
}
