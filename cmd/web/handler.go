package main

import "net/http"

func (app *application) homepage(w http.ResponseWriter, r *http.Request) {
	pageData := app.newTemplate(r)

	app.render(w, r, http.StatusOK, "homepage.gohtml", pageData)
}

func (app *application) createTemplateForm(w http.ResponseWriter, r *http.Request) {
	pageData := app.newTemplate(r)

	fontOptions := []string{"Roboto", "Open Sans", "Lato", "Montserrat", "Raleway"}
	pageData.FontOptions = fontOptions

	app.render(w, r, http.StatusOK, "template_form.gohtml", pageData)
}
