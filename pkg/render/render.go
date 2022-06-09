package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

// RenderTemplate renders template for fn handlers
func RenderTemplate(res http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(res, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}

func RenderTemplateTest(res http.ResponseWriter, tmpl string) error {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*page.tmpl")
	if err != nil {
		return err
	}

	for _, page := range pages {
		name := filepath.Base(page)
	}
}
