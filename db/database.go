package db

import (
    "database/sql"
	"log"
   _ "modernc.org/sqlite"
	
)

var DB *sql.DB

func InitDB(dataSourceName string)  {
    var err error
    DB, err = sql.Open("sqlite", dataSourceName)
    if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
}
