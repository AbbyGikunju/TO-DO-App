package models

import (
	"database/sql"
	"todo-app/db"
	"time"
)

type Todo struct {
	ID 				int
	UserID			int
	Title			string
	Priority		int
	DateCreated		time.Time
	Deadline        sql.NullTime
	Completed		bool
	Description 	sql.NullString
}

//Fetch todos for a given user
func GetTodosByUserID(userID int) ([]Todo, error) {
	rows, err := db.DB.Query(
		`SELECT id, user_id, title, priority, date_created, deadline, completed, description FROM todos WHERE user_id = ?`, userID) 
		if err != nil {
			return nil,  err
		}
		defer rows.Close()

		var todos []Todo
		for rows.Next() {
			var t Todo
			err := rows.Scan(
				&t.ID,
				&t.UserID,
				&t.Title,
				&t.Priority,
				&t.DateCreated,
				&t.Deadline,
				&t.Completed,
				&t.Description,
			)
			if err != nil {
				return nil, err
			}
			todos = append(todos, t)
		}
		return todos, nil
}