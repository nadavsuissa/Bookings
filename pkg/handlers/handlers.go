package handlers

import (
	"github.com/nadavsuissa/Bookings/pkg/config"
	"github.com/nadavsuissa/Bookings/pkg/models"
	"github.com/nadavsuissa/Bookings/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepo Creates a New Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers Sets the Reposritory for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteip := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteip)
	render.RenderTemplate(w, "home.page.gohtml", &models.TemplateData{})

}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"
	remoteip := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteip
	render.RenderTemplate(w, "about.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}
