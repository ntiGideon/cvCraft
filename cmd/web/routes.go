package main

import (
	"github.com/justinas/alice"
	"github.com/ntiGideon/cvCraft/ui"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// routes here
	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf)

	mux.Handle("GET /", dynamic.ThenFunc(app.homepage))
	mux.Handle("GET /create-template", dynamic.ThenFunc(app.createTemplateForm))
	mux.Handle("POST /create-template", dynamic.ThenFunc(app.createTemplate))
	mux.Handle("GET /list-templates", dynamic.ThenFunc(app.listTemplates))

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)
	return standard.Then(mux)
}
