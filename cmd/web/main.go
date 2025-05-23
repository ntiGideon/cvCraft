package main

import (
	"crypto/tls"
	"encoding/gob"
	"flag"
	"fmt"
	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	"github.com/joho/godotenv"
	"github.com/ntiGideon/cvCraft/cmd/service"
	"github.com/ntiGideon/cvCraft/ent"
	"github.com/ntiGideon/cvCraft/internal/dbClient"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type application struct {
	logger             *slog.Logger
	templateCache      map[string]*template.Template
	formDecoder        *form.Decoder
	sessionManager     *scs.SessionManager
	db                 *ent.Client
	templateRepository dbClient.TemplateRepository
	thumbnailService   service.ThumbnailService
}

func main() {
	gob.Register(map[string]interface{}{})

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))

	// postgres database connection
	connectionString := fmt.Sprintf("host=%v port=%v user=%v dbname=%s password=%s sslmode=%v", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASS"), os.Getenv("DB_SSL_MODE"))
	dbString := flag.String("dsn", connectionString, "database connection string")
	flag.Parse()
	db, err := connectDb(*dbString)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	dbDriver, err := OpenDriver(connectionString)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	formDecode := form.NewDecoder()

	// session manager
	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Secure = true
	sessionManager.Cookie.Persist = false
	sessionManager.Store = postgresstore.New(dbDriver)

	app := &application{
		logger:             logger,
		templateCache:      templateCache,
		formDecoder:        formDecode,
		db:                 db,
		sessionManager:     sessionManager,
		templateRepository: dbClient.TemplateRepository{Client: db},
		thumbnailService:   service.ThumbnailService{UploadPath: "ui/static/img/"},
	}

	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      app.routes(),
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Info("Starting server on port", "addr", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
