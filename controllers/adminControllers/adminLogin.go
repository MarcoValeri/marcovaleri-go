package admincontrollers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

type LoginValidation struct {
	EmailValidation    string
	PasswordValidation string
}

// Initialize the session
var store = sessions.NewCookieStore(securecookie.GenerateRandomKey(32), securecookie.GenerateRandomKey(32))

func AdminLogin() {
	tmpl := template.Must(template.ParseFiles("./views/admin/admin-login.html"))
	http.HandleFunc("/admin/login", func(w http.ResponseWriter, r *http.Request) {

		setLoginValidation := LoginValidation{
			EmailValidation:    "",
			PasswordValidation: "",
		}

		// Session authentication
		session, errSession := store.Get(r, "session-authentication")
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
			if getAdminUserEmail == "info@marcovaleri.net" && getAdminUserPassword == "1234" {
				session.Values["admin-user-authentication"] = true
				session.Save(r, w)
				http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
			} else {
				setLoginValidation.EmailValidation = "Error: email and password are not valid"
				setLoginValidation.PasswordValidation = "Error: email and password are not valid"
				session.Values["admin-user-authentication"] = false
				session.Save(r, w)
			}
		}

		tmpl.Execute(w, nil)
	})
}
