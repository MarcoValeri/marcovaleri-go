package admincontrollers

import (
	"fmt"
	"html/template"
	"marcovaleri/models"
	"marcovaleri/util"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

type LoginValidation struct {
	PageTitle          string
	EmailValidation    string
	PasswordValidation string
}

// Initialize the session
var store = sessions.NewCookieStore(securecookie.GenerateRandomKey(32), securecookie.GenerateRandomKey(32))

func AdminLogin() {
	tmpl := template.Must(template.ParseFiles("./views/admin/admin-login.html"))
	http.HandleFunc("/admin/login", func(w http.ResponseWriter, r *http.Request) {

		data := LoginValidation{
			PageTitle:          "Admin Login",
			EmailValidation:    "",
			PasswordValidation: "",
		}

		// Session authentication
		session, errSession := store.Get(r, "session-user-admin-authentication")
		if errSession != nil {
			fmt.Println("Error on session-authentication", errSession)
		}
		session.Values["admin-user-authentication"] = false
		session.Save(r, w)

		// Form validation
		getAdminUserEmail := r.FormValue("admin-user-email")
		getAdminUserPassword := r.FormValue("admin-user-password")
		getAdminUserLogin := r.FormValue("admin-user-login")

		if len(getAdminUserLogin) > 0 {

			// Email validation
			if !util.FormEmailInput(getAdminUserEmail) {
				data.EmailValidation = "Error: email format is not valid"
				session.Values["admin-user-authentication"] = false
				session.Save(r, w)
			}
			if !util.FormEmailLengthInput(getAdminUserEmail) {
				data.EmailValidation = "Error: email format is not valid"
				session.Values["admin-user-authentication"] = false
				session.Save(r, w)
			}

			// Password validation
			if !util.FormPasswordInput(getAdminUserPassword) {
				data.PasswordValidation = "Error: password is not valid"
				session.Values["admin-user-authentication"] = false
				session.Save(r, w)
			}

			// Form validation
			if models.UserAdminLogin(getAdminUserEmail, getAdminUserPassword) {
				session.Values["admin-user-authentication"] = true
				session.Save(r, w)
				http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
			} else {
				data.EmailValidation = "Error: email and password are not valid"
				data.PasswordValidation = "Error: email and password are not valid"
				session.Values["admin-user-authentication"] = false
				session.Save(r, w)
			}
		}

		tmpl.Execute(w, data)
	})
}
