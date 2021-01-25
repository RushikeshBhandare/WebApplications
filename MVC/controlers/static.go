package controlers

import (
	"WebApplications/MVC/Views"
)

func NewStactc() *Static {
	return &Static{
		Home:    Views.NewView("bootstrap", "Views/static/home.gohtml"),
		Contact: Views.NewView("bootstrap", "Views/static/contact.gohtml"),
	}
}

type Static struct {
	Home    *Views.View
	Contact *Views.View
}
