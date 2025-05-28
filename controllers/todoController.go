package controllers

import (
	"encoding/json"
	"net/http"
	

	"todo-app/utils"
	"todo-app/models"
)

//GetTaskHandler handles fetching tasks for the logged-in user

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := utils.GetUserIDFromSession(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	todos, err := models.GetTodosByUserID(userID)
	if err != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

