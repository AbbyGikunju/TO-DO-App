package models

import (
	"database/sql"
)

func CreateUser(db *sql.DB, username, email, passwordHash string) error {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, username, email, passwordHash)
	return err
} 

func GetUserByUsername(db *sql.DB, username string) (string, error) {
	var hashedPwd string
	query := `SELECT password FROM users WHERE username = ?`
	err := db.QueryRow(query, username).Scan(&hashedPwd)
	if err != nil {
		return "", err
	}
	return hashedPwd, nil
}
