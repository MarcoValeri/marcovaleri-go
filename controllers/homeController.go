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
	homeTemplate = template.Must(template.ParseFiles("./views/templates/base.html", "./views/home.html"))
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("./views/home.html")
	if err != nil {
		http.Error(w, "Error parsing temlate", http.StatusInternalServerError)
		log.Println("Template parsing error:", err)
	}

	data := PageData{
		PageTitle: "Marco Valeri",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		log.Println("Template execution error:", err)
		return
	}
}

// func Home() {
// 	tmpl :=
// 		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 			data := PageData{
// 				PageTitle: "Marco Valeri",
// 			}
// 			tmpl.Execute(w, data)
// 		})
// }
