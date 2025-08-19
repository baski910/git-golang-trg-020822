package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var templateCache = make(map[string]*template.Template)

func main() {
	// Build the template cache on startup
	if err := buildTemplateCache(); err != nil {
		log.Fatalf("Error building template cache: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get template from cache
		tmpl, ok := templateCache["home.html"]
		if !ok {
			http.Error(w, "Template not found", http.StatusInternalServerError)
			return
		}
		// Render the template
		tmpl.Execute(w, nil)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func buildTemplateCache() error {
	pages, err := filepath.Glob("./templates/*.html")
	if err != nil {
		return err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		tmpl, err := template.ParseFiles(page)
		if err != nil {
			return err
		}
		templateCache[name] = tmpl
	}
	return nil
}
