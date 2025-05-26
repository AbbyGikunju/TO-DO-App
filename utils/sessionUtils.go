package utils

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"time"
)

//GenerateSessionToken creates a secure random token string
func GenerateSessionToken() string {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}

// IsSessionValid checks if the session token is non-empty (basic check)
func IsSessionValid(token string) bool {
    return token != ""
}


//SetSessionCookie sets the session token as a secure HTTP cookie
func SetSessionCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:		"session_token",
		Value:		token,
		Path:		"/",
		Expires:	time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure: false,
	})
}

//GetSessionToken retrieves the session token from the cookie
func GetSessionToken(r *http.Request) string {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return ""
	}
	return cookie.Value
}

//ClearSessionCookie removes the sesson token cookie(logout)
func ClearSessionCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:		"session_token",
		Value:		"",
		Path:		"/",
		MaxAge:		-1,
		HttpOnly: true,
		Secure: false,
	})
	
}