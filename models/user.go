package models

import (
	"database/sql"
)

var DB *sql.DB

func SetDatabases(db *sql.DB) {
	DB = db
}

func CreateUser(username, email, passwordHash string) error {
	query := `INSERT INTO users (username, email, password) VALUES (?, ?, ?)`
	_, err := DB.Exec(query, username, email, passwordHash)
	return err
} 

func GetUserByUsername(username string) (string, error) {
	var hashedPwd string
	query := `SELECT password FROM user WHERE username = ?`
	err := DB.QueryRow(query, username).Scan(&hashedPwd)
	if err != nil {
		return "", err
	}
	return hashedPwd, nil
}