package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/streckc/go-bookings/internal/config"
	"github.com/streckc/go-bookings/internal/forms"
	"github.com/streckc/go-bookings/internal/models"
	"github.com/streckc/go-bookings/internal/render"
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
	render.RenderTemplate(w, r, "home.page.gohtml", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "about.page.gohtml", &models.TemplateData{})
}

// Majors renders the room page
func (m *Repository) MajorsRoom(w http.ResponseWriter, r *http.Request) {
	vs := map[string]string{
		"room_name": "Major's Suite",
	}
	render.RenderTemplate(w, r, "majors.page.gohtml", &models.TemplateData{
		StringMap: vs,
	})
}

// Generals renders the room page
func (m *Repository) GeneralsRoom(w http.ResponseWriter, r *http.Request) {
	vs := map[string]string{
		"room_name": "General's Quarters",
	}
	render.RenderTemplate(w, r, "generals.page.gohtml", &models.TemplateData{
		StringMap: vs,
	})
}

// Contact
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.gohtml", &models.TemplateData{})
}

// Availability renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.gohtml", &models.TemplateData{})
}

// PostAvailability renders the search availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte("posted to search availability for " + start + " to " + end))
	// render.RenderTemplate(w, r,"search-availability.page.gohtml", &models.TemplateData{})
}

// MakeReservation displays the make reservation form
func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation

	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.RenderTemplate(w, r, "make-reservation.page.gohtml", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// PostReservation handles the post data from make reservation
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("pohne"),
	}

	form := forms.New(r.PostForm)
	form.Required("first_name", "last_name", "email", "phone")
	form.MinLength("first_name", 3)
	form.MinLength("last_name", 2)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		render.RenderTemplate(w, r, "make-reservation.page.gohtml", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}
	log.Println(reservation)
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON handles rquest for avaialability and send JSON response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      false,
		Message: "Available!",
	}

	out, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
