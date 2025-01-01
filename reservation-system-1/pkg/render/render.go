package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ak-er/golang-playground/pkg/config"
	"github.com/ak-er/golang-playground/pkg/models"
)

// render template with more optimized version

var appconfig *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	appconfig = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var templateCache map[string]*template.Template
	if appconfig.UseCache {
		// get the template from the app config
		templateCache = appconfig.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCacheComplex()
	}

	// get request template cache
	t, isGet := templateCache[tmpl]

	if !isGet {
		log.Fatal("could not get the template from the templateCache")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td)
	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	// render template

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCacheComplex() (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}
	// get all file which name have *.page.tmpl from templates folder
	templateFiles, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return templateCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, templateFile := range templateFiles {
		templateName := filepath.Base(templateFile)
		templateSets, err := template.New(templateName).ParseFiles(templateFile)
		if err != nil {
			return templateCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil {
			return templateCache, err
		}

		if len(matches) > 0 {
			templateSets, err = templateSets.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return templateCache, err
			}
		}

		templateCache[templateName] = templateSets
	}
	return templateCache, nil
}
