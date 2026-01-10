package repository

import (
	"database/sql"

	"github.com/iamskyy666/golang-job-portal-api/internal/models"
)

func CreateJob(db *sql.DB, job *models.Job) (*models.Job, error) {
	result,err:=db.Exec(`INSERT INTO jobs (title, description, company, location, salary, user_id) VALUES (?, ?,?, ?, ?, ?)`,job.Title, job.Description, job.Company, job.Location, job.Salary, job.UserId)
	if err != nil {
		return nil,err
	}
	id,err:=result.LastInsertId()
	if err!=nil{
		return nil,err
	}

	job.ID = int(id)
	return job,nil
}