package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ak-er/golang-journey/pkg/config"
	"github.com/ak-er/golang-journey/pkg/models"
)

/*
	func RenderTemplate(w http.ResponseWriter, tmpl string) {
		// Render this without layout of template
		// parseTemplage, _ := template.ParseFiles("./templates/" + tmpl)
		// Render this with layout of template
		parseTemplage, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
		err := parseTemplage.Execute(w, nil)
		if err != nil {
			fmt.Println("error occur parsing template", err)
			return
		}
	}
*/

/*
// Render Template wth simple cache method
var templateCache = make(map[string]*template.Template)

	func RenderTemplate(w http.ResponseWriter, tmpl string) {
		var t *template.Template
		var err error

		// check if the template is already have in cache
		_, inCache := templateCache[tmpl]
		if !inCache {
			// need to create template
			err = createTemplateCache(tmpl)
			if err != nil {
				log.Println("Error to CreateTemplateCache:::", err)
			}
		}
		t = templateCache[tmpl]
		err = t.Execute(w, nil)
		if err != nil {
			log.Println("Error To ExecuteRenderTemplate:::", err)
		}
	}

	func createTemplateCache(tmpl string) error {
		templates := []string{
			fmt.Sprintf("./templates/%s", tmpl),
			"./templates/base.layout.tmpl",
		}

		// parse template
		t, err := template.ParseFiles(templates...)
		if err != nil {
			return err
		}

		// add template to cache
		templateCache[tmpl] = t
		return nil
	}
*/

/*

// render template with complex but better version for optimization
func RenderTemplate(w http.ResponseWriter, tmpl string) {

	// create template cache
	templateCache, err := CreateTemplateCacheComplex()
	if err != nil {
		log.Fatal(err)
	}

	// get request template cache
	t, isGet := templateCache[tmpl]

	if !isGet {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	err = t.Execute(buf, nil)
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

*/

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
