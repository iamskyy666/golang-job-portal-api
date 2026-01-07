package services

import (
	"database/sql"

	"github.com/iamskyy666/golang-job-portal-api/internal/models"
	"github.com/iamskyy666/golang-job-portal-api/internal/repository"
)

func GetUserById(db *sql.DB, id int) (*models.User,error) {
	return repository.GetUserById(db,id)
}