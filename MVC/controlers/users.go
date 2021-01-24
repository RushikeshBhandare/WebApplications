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

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	fmt.Println("USing newe methed")
	err := u.NewView.Render(w, nil)
	if err != nil {
		panic(err)
	}
}
