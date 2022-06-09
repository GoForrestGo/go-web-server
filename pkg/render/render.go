package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

/**
* RenderTemplate renders template for fn handlers
* Procedure:
  - Read Template from the disk
  - Parse the template
  - return it
*
*/

func RenderTemplate(res http.ResponseWriter, tmpl string) {
	// Just to show Template caching is working
	// _, err := RenderTemplateTest(res)

	// if err != nil {
	// 	fmt.Println("Error getting template cache", err)
	// }

	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// get the template cache from app config
	// ok is used (any second variable) with map to check if there's actually value in the map for
	// the corresponding key (return true) and (return false) in case of not found [ convernsion]
	t, ok := tc[tmpl]

	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer) // put the template into some bytes [buffer]

	_ = t.Execute(buf, nil) //take that t (template) and store its content into the buffer, nil (for not passign any data)
	_, err = buf.WriteTo(res)
	if err != nil {
		fmt.Println("Error Writing template to browser", err)
	}
	// parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	// err = parsedTemplate.Execute(res, nil)
	// if err != nil {
	// 	fmt.Println("error parsing template", err)
	// 	return
	// }
}

/**
* CreateTemplateCache: Parse all of the templates and store them into a map of templates "cache"
*
**/
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{} // creating a template cache that contains templates

	pages, err := filepath.Glob("./templates/*.page.tmpl") // get all the pages in that folder
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page) // to extract just the name of the page, not the full path
		fmt.Println("Page is currently", page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			// parse those templates
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}

// [about.page.tmpl] => full parsed and ready to use templates
// [home.page.tmpl] =>  full parsed and ready to use templates
