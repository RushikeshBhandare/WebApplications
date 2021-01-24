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

type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

//Create is used to prossess the sign up form wen a user
//Submit it .this is used rto creat a new user accout
//
//POST / signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	err := parseForm(r, &form)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, form)
}
