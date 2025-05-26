package routes

import (
	"database/sql"
	"todo-app/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router, db *sql.DB) {
	
	router.HandleFunc("/", controllers.LoginPageHandler).Methods("GET")
	router.HandleFunc("/index", controllers.IndexPageHandler).Methods("GET") 

	router.HandleFunc("/login", controllers.LoginHandler(db)).Methods("POST")
	router.HandleFunc("/register", controllers.RegisterPageHandler).Methods("GET")
	router.HandleFunc("/register", controllers.RegisterHandler(db)).Methods("POST")
	router.HandleFunc("/logout", controllers.LogoutHandler).Methods("GET") // Usually logout is a GET to clear cookie and redirect
}