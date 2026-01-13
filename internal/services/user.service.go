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

func UpdateProfilePicture(db *sql.DB,id int, profilePicture string)error{
	return repository.UpdateProfilePic(db,id,profilePicture)
}

func GetUsersService(db *sql.DB)([]*models.User,error){
	return repository.GetUsersRepo(db)
}