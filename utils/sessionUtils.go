package utils

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"errors"
	"net/http"
	"time"

	"todo-app/db"
)

//GetUserIDFromsession retrieves the user ID associated with the session token
func GetUserIDFromSession(r *http.Request) (int, error) {
	token := GetSessionToken(r)
	if token == "" {
		return 0, errors.New("missing session token")
	}
	var userID int
	err := db.DB.QueryRow("SELECT user_id FROM sessions where token = ?", token).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("invalid session token")
		}
		return 0, err
	}
	return userID, nil
}


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