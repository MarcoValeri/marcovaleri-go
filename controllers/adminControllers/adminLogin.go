package admincontrollers

import (
	"fmt"
	"html/template"
	"marcovaleri/models"
	"marcovaleri/session"
	"marcovaleri/util"
	"net"
	"net/http"
	"strings"
)

type LoginValidation struct {
	PageTitle          string
	EmailValidation    string
	PasswordValidation string
}

func getUserIpAddress(req *http.Request) string {
	userIps := req.Header.Get("X-Forwarded-For")

	// If X-Forwarded-For is not present, use the RemoteAddr
	if userIps == "" {
		// Extract the IP part from RemoteAddr (e.g., "192.0.2.1:12345" -> "192.0.2.1")
		ip, _, err := net.SplitHostPort(req.RemoteAddr)
		if err != nil {
			return "" // Or handle the error as appropriate
		}
		return ip
	}

	// Split the comma-separated IPs and trim spaces
	ips := strings.Split(userIps, ",")
	for i, ip := range ips {
		ips[i] = strings.TrimSpace(ip)
	}

	// Get the first IP from the list, as it's most likely the original client IP
	if len(ips) > 0 {
		return ips[0]
	}

	return "" // No valid IP found
}

func AdminLogin(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./views/admin/admin-login.html"))
	data := LoginValidation{
		PageTitle:          "Admin Login",
		EmailValidation:    "",
		PasswordValidation: "",
	}

	// Redirect IPs banned
	userIP := getUserIpAddress(r)
	isThisIpBanned, _ := models.UserAdminBannedByIp(userIP)
	if isThisIpBanned {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	if r.Method == http.MethodPost {
		// Session authentication
		session, errSession := session.Store.Get(r, "session-user-admin-authentication")
		if errSession != nil {
			fmt.Println("Error AdminLogin on session-authentication", errSession)
		}

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
				// Store user ip to the db
				models.UserAdminLoginIp(userIP, getAdminUserEmail, getAdminUserPassword)
				data.EmailValidation = "Error: email and password are not valid"
				data.PasswordValidation = "Error: email and password are not valid"
				session.Values["admin-user-authentication"] = false
				session.Save(r, w)
			}
		}
	}

	tmpl.Execute(w, data)
}
