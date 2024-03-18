// entry point to the front end of our application
package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

// a struct to hold the configuration information for our application
type config struct {
	port int    // port the application will listen on
	env  string // production/development
	api  string // backend url
	db   struct {
		dsn string // data source name: url to connect to the database
	}
	stripe struct {
		secret string // secret key
		key    string // publishable key
	}
}

// will use this struct as receiver for various functions of the application
type application struct {
	config        config
	infoLog       *log.Logger
	erroLog       *log.Logger
	templateCache map[string]*template.Template
	version       string
}

func (app *application) serve() error {
	// creating a http server
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port), // the address on which the server will listen on
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	app.infoLog.Printf("Starting HTTP server in %s mode on port %d", app.config.env, app.config.port)

	return srv.ListenAndServe()
}

func main() {
	// setting up the configuration of our application
	var cfg config

	// getting the configuration from flags set while running the application {--port|--env|--api}
	flag.IntVar(&cfg.port, "port", 4000, "Port on which the server listens on")
	flag.StringVar(&cfg.env, "env", "development", "Environment in which the application is running in {development|production}")
	flag.StringVar(&cfg.api, "api", "http://localhost:4001", "URL of the backend API service")

	flag.Parse()

	// The stripe secret key and publishable key should not be sent through command-line flags
	// An anthorised person can easily read them by doing ps -aux in the server
	// So, these will be read through environment variables
	cfg.stripe.secret = os.Getenv("ENV_STRIPE_SECRET")
	cfg.stripe.key = os.Getenv("ENV_STRIPE_KEY")

	// setting up our loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	templateCache := make(map[string]*template.Template)

	// setting up the application struct
	app := &application{
		config:        cfg,
		infoLog:       infoLog,
		erroLog:       errorLog,
		templateCache: templateCache,
		version:       version,
	}

	err := app.serve()
	if err != nil {
		app.erroLog.Println(err)
		log.Fatal(err)
	}
}
