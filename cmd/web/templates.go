package main

import (
	"github.com/ntiGideon/cvCraft/ui"
	"html/template"
	"io/fs"
	"path/filepath"
)

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// Regular pages
	pages, err := fs.Glob(ui.Files, "html/pages/*.gohtml")
	if err != nil {
		return nil, err
	}

	// CV templates
	cvTemplates, err := fs.Glob(ui.Files, "html/cv-templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	// Process regular pages
	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			"html/base.gohtml",
			"html/partials/*.gohtml",
			page,
		}

		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	// Process CV templates
	for _, templateFile := range cvTemplates {
		name := filepath.Base(templateFile)

		patterns := []string{
			"html/base.gohtml",
			"html/partials/*.gohtml",
			"html/cv-templates/partials/*.gohtml",
			templateFile,
		}

		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}

var functions = template.FuncMap{}
