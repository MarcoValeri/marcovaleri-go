package controllers

import (
	admincontrollers "marcovaleri/controllers/adminControllers"
	"net/http"
)

// func AdminController() {
// 	admincontrollers.AdminLogin()
// 	admincontrollers.AdminDashboard()
// 	admincontrollers.AdminUsers()
// 	admincontrollers.AdminUserAdd()
// }

func AdminController(w http.ResponseWriter, r *http.Request) {
	// This could be a router or a simple switch statement to handle different admin routes
	switch r.URL.Path {
	case "/admin/dashboard":
		// AdminDashboard(w, r)
		admincontrollers.AdminDashboard(w, r)
	// Add more cases for other admin routes like "/admin/users", "/admin/settings", etc.
	default:
		http.NotFound(w, r)
	}
}
