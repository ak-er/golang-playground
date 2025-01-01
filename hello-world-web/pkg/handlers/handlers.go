package handlers

import (
	"errors"
	"net/http"

	"github.com/ak-er/golang-journey/pkg/config"
	"github.com/ak-er/golang-journey/pkg/models"
	"github.com/ak-er/golang-journey/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

/*
func Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "This is the homepage")
	render.RenderTemplate(w, "home.page.tmpl")
}
*/

// optimized the cache
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr // get the remote_ip/address from request
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

/*
func About(w http.ResponseWriter, r *http.Request) {
	_, err := divideValue(3, 2)
	if err != nil {
		fmt.Fprintf(w, "%s", fmt.Sprintf("Something went wrong: %s", err))
		return
	}
	// _, _ = fmt.Fprintf(w, "%s", fmt.Sprintf("This is the about page and has sum calculation of %f", val))

	render.RenderTemplate(w, "about.page.tmpl", &render.TemplateData{})

}
*/

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Testing Text"
	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}

func divideValue(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("ZeroDivisionError: division by zero")
		return -1, err
	}
	return x / y, nil
}
