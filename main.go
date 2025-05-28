package main

import (
	"database/sql"
	"log"
	"net/http"

	"todo-app/db"
	"todo-app/routes"

	"github.com/gorilla/mux"
	 _ "modernc.org/sqlite"
)

func main() {
	//Initialize DB globally
	db.InitDB("todo.db")

	router := mux.NewRouter()
	db, err := sql.Open("sqlite", "./db/todo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // Close DB when main exits

	//Check if DB connection is alive
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	
	routes.SetupRoutes(router, db)

	log.Println("Starting server on: 8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
