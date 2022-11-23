package handlers

import (
	"net/http"

	"github.com/streckc/go-bookings/pkg/config"
	"github.com/streckc/go-bookings/pkg/models"
	"github.com/streckc/go-bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// Hone is the about page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	//log.Printf("%+v", r.RequestURI)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "helloooooo"
	stringMap["remote_ip"] = m.App.Session.GetString(r.Context(), "remote_ip")

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}
