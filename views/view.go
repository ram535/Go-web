package views

import (
	"path/filepath"
	"text/template"
)

var (
	layoutDir   = "views/layouts/"
	templateExt = ".gohtml"
)

func NewView(layout string, files ...string) *View {
	// ... unpack the slice, it is like it is passing one by one
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}

// Since we are not exporting this function,
// we not put a capital letter in the first letter
// of the name of the function
// layoutFiles return a slice of strings representing
// the layout files used in our application.
func layoutFiles() []string {
	files, err := filepath.Glob(layoutDir + "*" + templateExt)
	if err != nil {
		panic(err)
	}
	return files
}
