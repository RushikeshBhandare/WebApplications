package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome To My site</h1>")
	s := []string{"Sunny", "sanjay", "Sagar"}
	fmt.Fprint(w, "<ul>")
	for i := 0; i < len(s); i++ {
		fmt.Fprint(w, "<li>", s[i], "</li>")

	}
	fmt.Fprint(w, "</ul>")

	fmt.Println(w)

}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
