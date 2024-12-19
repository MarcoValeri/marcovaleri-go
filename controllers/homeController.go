package controllers

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	PageTitle string
}

// Cache the template
var homeTemplate *template.Template

// Parse the template once at startup
func init() {
	var err error
	homeTemplate, err = template.ParseFiles("./views/templates/base.html", "./views/home.html")
	if err != nil {
		log.Fatal("Error parsing template:", err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl := homeTemplate

	data := PageData{
		PageTitle: "Marco Valeri",
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		log.Fatal("Error parsing template:", err)
	}
}

// func NotFound(w http.ResponseWriter, r *http.Request) {
// 	renderErrorPage(w, http.StatusNotFound, "Page not found")
// }

// func renderErrorPage(w http.ResponseWriter, statusCode int, message string) {
// 	w.WriteHeader(statusCode)
// 	data := PageData{
// 		PageTitle: "Error",
// 	}
// 	err := errorTemplate.Execute(w, data)
// 	if err != nil {
// 		// Fallback to http.Error in case of error rendering the error page
// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
// 	}
// }
