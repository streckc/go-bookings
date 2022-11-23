package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/streckc/go-bookings/pkg/config"
	"github.com/streckc/go-bookings/pkg/handlers"
	"github.com/streckc/go-bookings/pkg/render"
)

const portNumber = ":8000"

var app config.AppConfig

var session *scs.SessionManager

func init() {
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create the template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false
	app.InProduction = false

	render.NewTemplates(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session
}

// main is the main entry point of the application
func main() {
	log.Printf("Starting server on %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err := srv.ListenAndServe()
	log.Fatal(err)
}
