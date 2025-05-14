package data

import (
	"todo-app/utils"
	"todo-app/models"
)

func UserIsValid(uName, pwd string) bool {
	hashedPwd, err := models.GetUserByUsername(uName)
	if err != nil {
		return false
	}
	return utils.VerifyPassword(pwd, hashedPwd)
}