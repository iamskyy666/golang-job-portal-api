package services

import (
	"database/sql"

	"github.com/iamskyy666/golang-job-portal-api/internal/models"
	"github.com/iamskyy666/golang-job-portal-api/internal/repository"
)

func CreateJob(db *sql.DB, job *models.Job)(*models.Job,error) {
	return repository.CreateJob(db,job)
}