package data

import (
	"database/sql"
	"todo-app/utils"
	"todo-app/models"
)

func UserIsValid(db *sql.DB, uName, pwd string) bool {
	hashedPwd, err := models.GetUserByUsername(db, uName)
	if err != nil {
		return false
	}
	return utils.VerifyPassword(pwd, hashedPwd)
}
