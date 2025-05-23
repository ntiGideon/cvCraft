package main

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"github.com/justinas/nosurf"
	"github.com/ntiGideon/cvCraft/ent"
	"github.com/ntiGideon/cvCraft/ent/migrate"
	"time"

	//"github.com/ntiGideon/cvCraft/ent"
	//"github.com/ntiGideon/cvCraft/ent/migrate"
	"github.com/ntiGideon/cvCraft/internal/models"
	"net/http"
	"runtime/debug"

	_ "github.com/lib/pq"
)

type templateData struct {
	Form        any
	Flash       string
	Toast       map[string]interface{}
	CV          ent.Resume
	Config      models.TemplateConfig
	CurrentYear int
	CSRFToken   string
	FontOptions any

	Templates ent.Templates
	IsAdmin   bool
	HasMore   bool
	NextPage  int
}

func connectDb(connectionString string) (*ent.Client, error) {
	client, err := ent.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	err = client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true))
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected to the database")
	return client, nil
}

func OpenDriver(connectionString string) (*sql.DB, error) {
	drv, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return drv, nil
}

func (app *application) newTemplate(r *http.Request) templateData {
	return templateData{
		CSRFToken:   nosurf.Token(r),
		CurrentYear: time.Now().Year(),
		Flash:       app.sessionManager.PopString(r.Context(), "flash"),
	}
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		trace  = string(debug.Stack())
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri, "trace", trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("<h1>An Error occurred at our side!</h1>"))
	app.logger.Error("An error occurred at our side!", "status", status, "trace", string(debug.Stack()))
	http.Error(w, http.StatusText(status), status)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data templateData) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, r, err)
		return
	}
	buf := new(bytes.Buffer)
	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	w.WriteHeader(status)

	_, err = buf.WriteTo(w)
	if err != nil {
		return
	}
}
