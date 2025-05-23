package main

import (
	"bytes"
	"fmt"
	"github.com/ntiGideon/cvCraft/ent"

	//"github.com/ntiGideon/cvCraft/ent"
	"github.com/ntiGideon/cvCraft/internal/models"
	"net/http"
)

type CVTemplateData struct {
	CV     *ent.Resume
	Config *models.TemplateConfig
}

func (app *application) renderCV(w http.ResponseWriter, r *http.Request, cv *ent.Resume, tmpl *ent.Template) {
	// Determine which template to use based on the template style
	templateName := "cv_" + tmpl.Style + ".gohtml" // e.g., "cv_classic.gohtml"

	// Prepare template data
	data := templateData{
		CV:     *cv,
		Config: *tmpl.Config,
	}

	// Use your existing render function
	app.render(w, r, http.StatusOK, templateName, data)
}

// If you need to render to a string (for PDF generation etc.)
func (app *application) renderCVToString(r *http.Request, cv *ent.Resume, tmpl *ent.Template) (string, error) {
	templateName := "cv_" + tmpl.Style + ".gohtml"
	ts, ok := app.templateCache[templateName]
	if !ok {
		return "", fmt.Errorf("template %s not found", templateName)
	}

	data := templateData{
		CV:     *cv,
		Config: *tmpl.Config,
	}

	var buf bytes.Buffer
	err := ts.ExecuteTemplate(&buf, "base", data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
