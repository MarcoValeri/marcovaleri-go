package main

import (
	"marcovaleri/controllers"
	"net/http"
)

func main() {
	// Static files
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	// Controllers
	controllers.Home()
	controllers.AdminController()

	http.ListenAndServe(":80", nil)
}
