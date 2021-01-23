package Views

import (
	"html/template"
)

func NewView(Layout string, files ...string) *View {
	files = append(files,
		"Views/layouts/footer.gohtml",
		"Views/layouts/bootstrap.gohtml",
		"Views/layouts/Navbar.gohtml",
	)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   Layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}
