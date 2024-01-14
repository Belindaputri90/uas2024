package handlers

import (
    "net/http"
    "strings"

    "github.com/dgrijalva/jwt-go"
)

// Check if the user is logged in
func IsUserLoggedIn(r *http.Request) bool {
    cookie, err := r.Cookie("session_token")
    return err == nil && isValidSessionToken(cookie.Value)
}

// Check if the session token is valid
func isValidSessionToken(token string) bool {
    if token == "" {
        return false
    }

    parts := strings.Split(token, " ")
    if len(parts) != 2 {
        return false
    }

    // Verifikasi struktur token JWT
    _, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
        return []byte(globalSecretKey), nil
    })

    return err == nil
}
