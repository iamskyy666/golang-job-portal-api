package services

import (
	"database/sql"

	"github.com/iamskyy666/golang-job-portal-api/internal/models"
	"github.com/iamskyy666/golang-job-portal-api/internal/repository"
)

func GetUserById(db *sql.DB, id int) (*models.User,error) {
	return repository.GetUserById(db,id)
}

func UpdateUserProfile(db *sql.DB, id int, username, emailId string)(*models.User,error){
	user:=&models.User{ID:id,Username: username, Email: emailId}
	return repository.UpdateUser(db,user)
}