package controllers

import (
	"html/template"
	"net/http"
)

type PageData struct {
	PageTitle string
}

func Home() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/home.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			PageTitle: "Marco Valeri",
		}
		tmpl.Execute(w, data)
	})
}
