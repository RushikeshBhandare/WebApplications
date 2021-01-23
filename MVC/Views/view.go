package Views

import (
	"html/template"
	"path/filepath"
)

var (
	LayoutDir string = "Views/layouts/"
	Extention string = ".gohtml"
)

func NewView(Layout string, files ...string) *View {
	files = append(files, layoutFiles()...)

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

//layout files return a slice of strings representing
//the layout files used in our application
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + Extention)
	if err != nil {
		panic(err)
	}
	return files
}
