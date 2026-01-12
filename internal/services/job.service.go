package services

import (
	"database/sql"
	"errors"

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


func UpdateJobService(db *sql.DB, job *models.Job,userID int, isAdmin bool)(*models.Job,error){
	existingJob,err:= repository.GetJobByIdRepo(db, job.ID)
	if err!=nil{
		return nil,err
	}

	if !isAdmin && existingJob.UserId != userID{
		return nil, errors.New("unauthorized to update this job!")
	}

	return repository.UpdateJobRepo(db,job)
}


func DeleteJobService(db *sql.DB, id int,userID int, isAdmin bool)error{
	existingJob,err:= repository.GetJobByIdRepo(db, id)
	if err!=nil{
		return err
	}

	if !isAdmin && existingJob.UserId != userID{
		return errors.New("unauthorized to DELETE this job!")
	}

	return repository.DeleteJobRepo(db,id)
}
