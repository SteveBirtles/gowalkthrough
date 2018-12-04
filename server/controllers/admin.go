package controllers

import (
	"../../server/models"
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"
)

func AdminLogin(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Println("/admin/login/", username)

	allAdmins := models.SelectAllAdmins()

	for _, a := range allAdmins {

		if a.Username == username {
			if a.Password != password {
				_, _ = fmt.Fprintln(w, "Error: Incorrect password")
				return
			}
			token, _ := uuid.NewV4()
			a.SessionToken = token.String()
			models.UpdateAdmin(a)
			_, _ = fmt.Fprintln(w, a.SessionToken)
		}

	}

}

func AdminCheck(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("sessionToken")
	if err != nil {
		fmt.Println("Error: No session cookie provided")
		return
	}

	fmt.Println("/admin/check/", cookie.Value)

	username := ValidateSessionToken(cookie.Value)
	_, _ = fmt.Fprintln(w, username)

}

func ValidateSessionToken(token string) string {
	allAdmins := models.SelectAllAdmins()
	for _, a := range allAdmins {
		if a.SessionToken == token {
			return a.Username
		}
	}
	return ""
}
