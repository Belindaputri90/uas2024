// handlers/auth.go
package handlers

import (
	"net/http"
)

// Check if the user is logged in
func IsUserLoggedIn(r *http.Request) bool {
	cookie, err := r.Cookie("session_token")
	return err == nil && isValidSessionToken(cookie.Value)
}

// Check if the session token is valid
func isValidSessionToken(token string) bool {
	return token != ""
}
