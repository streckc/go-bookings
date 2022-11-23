package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templateCache = make(map[string]*template.Template)

// RenderTemplate parses the template and outputs to writer
func RenderTemplateTest(w http.ResponseWriter, t string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+t, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Printf("error parsing template: %v", err)
		return
	}
}

// RenderTemplate parses the template and outputs to writer
func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check to see if we already have the template in our cache
	_, inMap := templateCache[t]
	if !inMap {
		// create template
		log.Printf("creating template: %s", t)
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// we have template in cache
		log.Printf("using cached template: %s", t)
	}

	tmpl = templateCache[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	templateCache[t] = tmpl
	return nil
}
