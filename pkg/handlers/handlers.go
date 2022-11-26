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
	// if r.URL.Path != "/" {
	// 	http.NotFound(w, r)
	// 	return
	// }
	render.RenderTemplate(w, "home.page.gohtml", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.gohtml", &models.TemplateData{})
}

// Majors
func (m *Repository) MajorsRoom(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "majors.page.gohtml", &models.TemplateData{})
}

// Generals
func (m *Repository) GeneralsRoom(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "generals.page.gohtml", &models.TemplateData{})
}

// Contact
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.gohtml", &models.TemplateData{})
}

// Reservation
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "reservation.page.gohtml", &models.TemplateData{})
}
