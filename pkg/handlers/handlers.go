package handlers

import (
	"net/http"

	"github.com/ankitkr1924/BankManagementApp/pkg/config"
	"github.com/ankitkr1924/BankManagementApp/pkg/models"
	"github.com/ankitkr1924/BankManagementApp/pkg/render"
)

var Repo *Repositry

type Repositry struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repositry {
	return &Repositry{
		App: a,
	}
}

func NewHandlers(r *Repositry) {
	Repo = r
}

func (m *Repositry) Home(rw http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(rw, "home.page.tmpl", &models.TemplateData{})

}

func (m *Repositry) About(rw http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIp
	render.RenderTemplate(rw, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
