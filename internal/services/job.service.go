package services

import (
	"database/sql"

	"github.com/iamskyy666/golang-job-portal-api/internal/models"
	"github.com/iamskyy666/golang-job-portal-api/internal/repository"
)

func CreateJob(db *sql.DB, job *models.Job)(*models.Job,error) {
	return repository.CreateJob(db,job)
}

func GetAllJobsService(db *sql.DB)([]models.Job,error){
	return repository.GetAllJobsRepo(db)
}

func GetJobsByUserIdService(db *sql.DB, userID int)([]models.Job,error){
	return repository.GetJobsByUserIdRepo(db, userID)
}

func GetJobByIdService(db *sql.DB, id int)(*models.Job,error){
	return repository.GetJobByIdRepo(db, id)
}