package controllers

import (
	//"todo-app/data"
	"fmt"
	"net/http"
	"todo-app/helpers"
	"todo-app/models"
	"todo-app/utils"

	//"github.com/gorilla/securecookie"
)

// var cookieHandler = securecookie.New(
// 	securecookie.GenerateRandomKey(64),
// 	securecookie.GenerateRandomKey(32),
// )

//Handlers
//Function to register users
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	uName := r.FormValue("username")
	email := r.FormValue("email")
	pwd := r.FormValue("password")
	confirmpwd := r.FormValue("confirmPassword")

	if helpers.IsEmpty(uName) || helpers.IsEmpty(email) || helpers.IsEmpty(pwd) || helpers.IsEmpty(confirmpwd) {
		fmt.Fprintln(w, "Fields cannot be empty!")
		return
	}
	if pwd != confirmpwd {
		fmt.Fprintln(w, "Passwords do not match")
		return
	}

	hashedPwd, err := utils.HashPassword(pwd)
	if err != nil {
		fmt.Fprintln(w, "Error hashing password:", err)
		return
	}

	err = models.CreateUser(uName, email, hashedPwd)
	if err != nil {
		fmt.Fprintln(w, "Failed to register user:", err)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}