package controlers

import (
	"WebApplications/MVC/Views"
	"fmt"
	"net/http"
)

//new users use to create a new user conmtroller
//this function can panic if thr template are not
//parsed correctolly and shouls onny used during intial setup
func NewUser() *Users {
	return &Users{
		NewView: Views.NewView("bootstrap", "Views/users/new.gohtml"),
	}
}

type Users struct {
	NewView *Views.View
}

//New is used to render ther form hwen a user can
//create new used account
//
//GET / SIGNUP
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	err := u.NewView.Render(w, nil)
	if err != nil {
		panic(err)
	}
}

//Create is used to prossess the sign up form wen a user
//Submit it .this is used rto creat a new user accout
//
//POST / signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	// r.PostForm
	fmt.Fprint(w, r.PostFormValue("email"))
	fmt.Fprint(w, r.PostForm["email"])

	fmt.Fprint(w, r.PostFormValue("password"))
	fmt.Fprint(w, r.PostForm["password"])

	fmt.Fprint(w, "this is temprarhy response")

}
