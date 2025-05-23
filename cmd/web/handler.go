package main

import (
	"github.com/ntiGideon/cvCraft/internal/models"
	"github.com/ntiGideon/cvCraft/internal/validator"
	"net/http"
)

func (app *application) getToast(r *http.Request) map[string]interface{} {
	val := app.sessionManager.Pop(r.Context(), "toast")
	if val == nil {
		return nil
	}

	toast, ok := val.(map[string]interface{})
	if !ok {
		app.logger.Error("could not convert toast toast")
		return nil
	}

	if toast["Message"] == nil {
		return nil
	}

	return toast
}

func (app *application) homepage(w http.ResponseWriter, r *http.Request) {
	pageData := app.newTemplate(r)
	pageData.Toast = app.getToast(r)

	app.render(w, r, http.StatusOK, "homepage.gohtml", pageData)
}

func (app *application) createTemplateForm(w http.ResponseWriter, r *http.Request) {
	pageData := app.newTemplate(r)

	fontOptions := []string{"Roboto", "Open Sans", "Lato", "Montserrat", "Raleway"}
	pageData.FontOptions = fontOptions

	app.render(w, r, http.StatusOK, "template_form.gohtml", pageData)
}

func (app *application) createTemplate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	dto := models.CreateTemplateConfigDto{
		Name:        r.PostForm.Get("name"),
		Description: r.PostForm.Get("description"),
		Premium:     r.PostForm.Get("premium") == "on",
		Config: models.TemplateConfig{
			ColorScheme: models.ColorScheme{
				Primary:    r.FormValue("primary_color"),
				Secondary:  r.FormValue("secondary_color"),
				Background: r.FormValue("background_color"),
				Text:       r.FormValue("text_color"),
				Accent:     r.FormValue("accent_color"),
				Header:     r.FormValue("header_color"),
			},
			FontSettings: models.FontSettings{
				HeaderFont: r.FormValue("header_font"),
				BodyFont:   r.FormValue("body_font"),
				FontSize:   r.FormValue("font_size"),
				LineHeight: r.FormValue("line_height"),
			},
			SectionOrder: []string{
				"personal_info",
				"summary",
				"experience",
				"education",
				"skills",
			},
			LayoutOptions: models.LayoutOptions{
				Columns: func() int {
					if r.FormValue("columns") == "2" {
						return 2
					} else {
						return 1
					}
				}(),
				SectionSpacing: r.FormValue("section_spacing"),
				HeaderStyle:    r.FormValue("header_style"),
			},
		},
	}
	dto.Config.Name = dto.Name

	dto.CheckField(validator.NotBlank(dto.Name), "name", "This field is required")

	if !dto.Valid() {
		data := app.newTemplate(r)
		data.Form = dto
		app.render(w, r, http.StatusUnprocessableEntity, "template_form.gohtml", data)
		return
	}

	thumbnailUrl, err := app.thumbnailService.GenerateFromConfig(&dto.Config)
	if err != nil {
		app.logger.Error("failed to generate thumbnail", "error", err)
		app.clientError(w, http.StatusBadRequest)
		return
	}
	dto.Thumbnail = thumbnailUrl

	err = app.templateRepository.CreateTemplate(r.Context(), &dto)
	if err != nil {
		app.logger.Error("failed to create template", "error", err.Error())
		app.serverError(w, r, err)
		return
	}

	toastDto := map[string]interface{}{
		"Type":    "success",
		"Message": "Resume template successfully created!",
	}
	app.sessionManager.Put(r.Context(), "toast", toastDto)
	http.Redirect(w, r, "/list-templates", http.StatusSeeOther)
}

func (app *application) listTemplates(w http.ResponseWriter, r *http.Request) {
	pageData := app.newTemplate(r)

	availableTemplates, err := app.templateRepository.ListTemplates(r.Context())
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Check if user is admin (for edit/create buttons)
	//isAdmin := app.isAuthenticated(r) && app.isAdmin(r)
	pageData.IsAdmin = true

	pageData.Templates = availableTemplates
	pageData.Toast = app.getToast(r)
	app.render(w, r, http.StatusOK, "list-templates.gohtml", pageData)
}
