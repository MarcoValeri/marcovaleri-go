package main

import (
	"marcovaleri/controllers"
	"net/http"
)

func main() {
	// Static files
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	mux := http.NewServeMux()

	// Controllers
	// controllers.Home()
	controllers.AdminController()

	mux.HandleFunc("/", controllers.Home)

	http.ListenAndServe(":80", mux)
}
