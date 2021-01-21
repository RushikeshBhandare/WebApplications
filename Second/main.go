package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcoem to my page</h1>")
	} else if r.URL.Path == "/Contact" {
		fmt.Fprint(w, "To get in touch please send an email to <a href=\"google.com\" >Gmail</a>")
	}

}

func main() {
	http.HandleFunc("/", handleFunc)
	http.ListenAndServe(":3000", nil)
}
