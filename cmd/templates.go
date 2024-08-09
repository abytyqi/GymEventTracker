package main

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.New("").ParseGlob("templates/**/*.html"))
}

func renderTemplate(w http.ResponseWriter, tmplName string, data interface{}) error {
	return tmpl.ExecuteTemplate(w, tmplName, data)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	err := renderTemplate(w, "pages/home.html", map[string]interface{}{"Title": "Home"})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	err := renderTemplate(w, "pages/about.html", map[string]interface{}{"Title": "About"})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
