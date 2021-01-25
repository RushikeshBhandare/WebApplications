package main

import (
	"WebApplications/MVC/controlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func PageNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Error 404 Webpage not found</h1><p>Contact if redirectd to Anothe page </p>")
}

func main() {

	statcC := controlers.NewStactc()
	UserC := controlers.NewUser()

	r := mux.NewRouter()
	r.Handle("/", statcC.Home).Methods("GET")
	r.Handle("/Contact", statcC.Contact).Methods("GET")
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
