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

// func GetAllJobsRepo(db *sql.DB) ([]models.Job, error){
// rows,err:=db.Query("SELECT * FROM jobs")
// if err != nil {
// 	return nil, err
// }
// defer rows.Close()

// var jobs []models.Job
// for rows.Next(){
// 	var job models.Job
// 	if err:=rows.Scan(&job.ID, &job.Title, &job.Description, &job.Location, &job.Company, &job.Salary, &job.CreatedAt,&job.UserId);err!=nil{
// 		return nil,err
// 	}
// 	jobs = append(jobs, job)
//  }
// 	return jobs,nil
// }

func GetAllJobsRepo(db *sql.DB) ([]models.Job, error) {
	rows, err := db.Query(`
		SELECT 
			id,
			title,
			description,
			location,
			company,
			salary,
			created_at,
			user_id
		FROM jobs
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []models.Job

	for rows.Next() {
		var job models.Job
		if err := rows.Scan(
			&job.ID,
			&job.Title,
			&job.Description,
			&job.Location,
			&job.Company,
			&job.Salary,
			&job.CreatedAt,
			&job.UserId,
		); err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return jobs, nil
}


	//! The job Model {} for refernce..
	// type Job struct {
	// ID          int       `json:"id"`
	// Title       string    `json:"title"`
	// Description string    `json:"description"`
	// Location    string    `json:"location"`
	// Company     string    `json:"company"`
	// Salary		string    `json:"salary"`
	// CreatedAt   time.Time `json:"created_at"`
	// UserId 		int       `json:"user_id"`
	// }