package controllers

import (
	admincontrollers "marcovaleri/controllers/adminControllers"
	"net/http"
)

func AdminController(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/admin/dashboard":
		admincontrollers.AdminDashboard(w, r)
	case "/admin/admin-users":
		admincontrollers.AdminUsers(w, r)
	case "/admin/admin-user-add":
		admincontrollers.AdminUserAdd(w, r)
	default:
		http.NotFound(w, r)
	}
}
