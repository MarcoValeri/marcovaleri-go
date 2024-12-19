package main

import (
	"marcovaleri/controllers"
	admincontrollers "marcovaleri/controllers/adminControllers"
	"net/http"
)

func main() {
	// Static files
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	mux := http.NewServeMux()

	// Controllers
	// controllers.Home()
	// controllers.AdminController()

	// mux.Handle("/admin/", controllers.AuthMiddleware(http.HandlerFunc(controllers.AdminController)))
	mux.Handle("/admin/", controllers.AuthMiddlewareController(http.HandlerFunc(controllers.AdminController)))

	mux.HandleFunc("/", controllers.Home)
	mux.HandleFunc("/admin/login", admincontrollers.AdminLogin)

	http.ListenAndServe(":80", mux)
}
