package Views

import (
	"html/template"
	"net/http"
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

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := v.Render(w, nil)
	if err != nil {
		panic(err)
	}
}

//Render is used to under the view
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
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
