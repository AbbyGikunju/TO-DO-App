package controllers

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"todo-app/helpers"
	"todo-app/models"
	"todo-app/utils"

	"github.com/gorilla/securecookie"
	"golang.org/x/crypto/bcrypt"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32),
)

// RegisterPageHandler renders the registration page
func RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/register.html")
	if err != nil {
		http.Error(w, "Unable to load registration page", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Failed to render registration page", http.StatusInternalServerError)
		return
	}
}

// RegisterHandler handles user registration
func RegisterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var req struct {
			Username        string `json:"username"`
			Email           string `json:"email"`
			Password        string `json:"password"`
			ConfirmPassword string `json:"confirmPassword"`
		}

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, `{"success": false, "message": "Invalid request body"}`, http.StatusBadRequest)
			return
		}

		if helpers.IsEmpty(req.Username) || helpers.IsEmpty(req.Email) || helpers.IsEmpty(req.Password) || helpers.IsEmpty(req.ConfirmPassword) {
			http.Error(w, `{"success": false, "message": "Fields cannot be empty!"}`, http.StatusBadRequest)
			return
		}

		if req.Password != req.ConfirmPassword {
			http.Error(w, `{"success": false, "message": "Passwords do not match"}`, http.StatusBadRequest)
			return
		}

		hashedPwd, err := utils.HashPassword(req.Password)
		if err != nil {
			http.Error(w, `{"success": false, "message": "Error hashing password"}`, http.StatusInternalServerError)
			return
		}

		err = models.CreateUser(db, req.Username, req.Email, hashedPwd)
		if err != nil {
			log.Println("Register error:", err)
			http.Error(w, `{"success": false, "message": "Failed to register user!"}`, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"success": true, "message": "User registered successfully"}`))
	}
}

// LoginPageHandler renders the login page
func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		http.Error(w, "Unable to load login page", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Failed to render login page", http.StatusInternalServerError)
		return
	}
}

// LoginRequest represents the expected login JSON payload
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginHandler handles user login
func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, `{"success": false, "message": "Invalid Request"}`, http.StatusBadRequest)
			return
		}

		var userID int
		var hashedPassword string
		err = db.QueryRow("SELECT id, password FROM users WHERE username = ?", req.Username).Scan(&userID, &hashedPassword)
		if err == sql.ErrNoRows {
			http.Error(w, `{"success": false, "message": "User not found"}`, http.StatusUnauthorized)
			return
		} else if err != nil {
			http.Error(w, `{"success": false, "message": "Database error"}`, http.StatusInternalServerError)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password))
		if err != nil {
			http.Error(w, `{"success": false, "message": "Incorrect password"}`, http.StatusUnauthorized)
			return
		}

		token := utils.GenerateSessionToken()
		utils.SetSessionCookie(w, token)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true}`))
	}
}


func IndexPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "")
}

// LogoutHandler logs the user out by clearing session cookie
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	utils.ClearSessionCookie(w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
