package admincontrollers

import (
	"fmt"
	"html/template"
	"net/http"
)

func AdminUsers() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-users.html"))
	http.HandleFunc("/admin/admin-users", func(w http.ResponseWriter, r *http.Request) {

		session, errSession := store.Get(r, "session-user-admin-authentication")
		if errSession != nil {
			fmt.Println("Error on session-authentication:", errSession)
		}

		if session.Values["admin-user-authentication"] == true {
			tmpl.Execute(w, nil)
		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}
	})
}

func AdminUserAdd() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-user-add.html"))
	http.HandleFunc("/admin/admin-user-add", func(w http.ResponseWriter, r *http.Request) {

		session, errSession := store.Get(r, "session-user-admin-authentication")
		if errSession != nil {
			fmt.Println("Error on session-authentication:", errSession)

			// Flag validation
			var areAdminUserInputsValid [5]bool
			isFormSubmittionValid := false

			// Get value from the form
			getAdminUserEmail := r.FormValue("admin-user-email")
			getAdminUserPassword := r.FormValue("admin-user-password")
			getAdminUserPasswordRepeat := r.FormValue("admin-user-password-repeat")
			getAdminUserSubmit := r.FormValue("admin-user-password-submit")
		}

		if session.Values["admin-user-authentication"] == true {
			tmpl.Execute(w, nil)
		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}
	})
}
