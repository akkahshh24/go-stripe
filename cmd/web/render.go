package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
)

// struct for passing information to the template
type templateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string // used for cross-site request forgery protection
	Flash           string // message which is shown once
	Warning         string
	Error           string
	isAuthenticated int
	API             string
	CSSVersion      string
}

var functions = template.FuncMap{}

//go:embed templates
var templateFS embed.FS // to compile our application with all of it's associated templates into a single binary

func (app *application) addDefaultData(td *templateData, r http.Request) *templateData {
	return td
}

func (app *application) renderTemplate(w http.ResponseWriter, r *http.Request, page string, td *templateData, partials ...string) {
	var (
		t   *template.Template
		err error
	)
	templateToRender := fmt.Sprintf("templates/%s/page.tmpl", page)

	// check if the template to render exists in the templateCache
	_, templateInMap := app.templateCache[templateToRender]
}
