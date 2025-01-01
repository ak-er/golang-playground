package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ak-er/golang-playground/pkg/config"
	"github.com/ak-er/golang-playground/pkg/handlers"
	"github.com/ak-er/golang-playground/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const PORT = "8080"

var session *scs.SessionManager

func main() {
	var appConfig config.AppConfig
	appConfig.InProduction = false // true when the project goes to production
	session = scs.New()
	session.Lifetime = 24 * time.Hour // set the expiration of session
	session.Cookie.Persist = true     // true when the browser close then stored before session
	session.Cookie.Secure = appConfig.InProduction
	session.Cookie.SameSite = http.SameSiteLaxMode

	appConfig.Session = session
	templateCache, err := render.CreateTemplateCacheComplex()

	appConfig.TemplateCache = templateCache
	appConfig.UseCache = false

	repo := handlers.NewRepo(&appConfig)
	handlers.NewHandlers(repo)

	render.NewTemplates(&appConfig)

	if err != nil {
		log.Fatal("error occur to CreateTemplateCache::", err)
	}
	// Routes()
	fmt.Println(fmt.Sprintf("Server is starting at PORT: %s", PORT))
	srv := &http.Server{
		Addr:    ":" + PORT,
		Handler: routesChi(&appConfig),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
