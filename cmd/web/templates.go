package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"time"

	"snippetbox.owezzy.tech/internal/models"
	"snippetbox.owezzy.tech/ui"
)

type templateData struct {
	CurrentYear     int
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
}

// Create a humanDate function which returns a nicely formatted string

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")

}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	// new mapto act as cache

	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		// Create a slice containing the filepath patterns for the templates we
		// want to parse.
		patterns := []string{
			"./ui/html/partials/*.html",
			"./ui/html/base.html",
		}
		// Use ParseFS() instead of ParseFiles() to parse the template files
		// from the ui.Files embedded filesystem.

		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		// Add the template set to the map, using the name of the page
		cache[name] = ts
	}
	// Return the map.
	return cache, nil
}
