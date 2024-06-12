package controllers

import admincontrollers "marcovaleri/controllers/adminControllers"

func AdminController() {
	admincontrollers.AdminLogin()
	admincontrollers.AdminDashboard()
	admincontrollers.AdminUsers()
	admincontrollers.AdminUserAdd()
}
