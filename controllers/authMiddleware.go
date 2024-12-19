package controllers

import (
	"fmt"
	"marcovaleri/session"
	"net/http"
)

func AuthMiddlewareController(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, errSession := session.Store.Get(r, "session-user-admin-authentication")
		if errSession != nil {
			fmt.Println("Error 1 on session-authentication:", errSession)
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther) // Redirect to login on error
			return
		}

		if session.Values["admin-user-authentication"] == true {
			// User is authenticated, proceed to the next handler
			next.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}
	})
}
