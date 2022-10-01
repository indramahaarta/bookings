package handlers

import (
	"net/http"

	"github.com/indramahaarta/bookings/pkg/config"
	"github.com/indramahaarta/bookings/pkg/models"
	"github.com/indramahaarta/bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	render.RenderTemplate(w, "home.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, Indra Mahaarta"
	stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(w, "about.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
