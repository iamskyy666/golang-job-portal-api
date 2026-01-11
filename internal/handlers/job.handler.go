package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iamskyy666/golang-job-portal-api/internal/models"
	"github.com/iamskyy666/golang-job-portal-api/internal/services"
)

func CreateJobHandler(db *sql.DB) gin.HandlerFunc{
	return func(ctx *gin.Context) {
		var job models.Job

		if err:=ctx.ShouldBindJSON(&job);err!=nil{
			ctx.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return 
		}

		userID:=ctx.GetInt("userID")
		job.UserId = userID

		createdJob,err:=services.CreateJob(db,&job);
		if err!=nil{
			ctx.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
			return 
		}

		ctx.JSON(http.StatusCreated, createdJob)
	}
}

func GetAllJobsHandler(db *sql.DB) gin.HandlerFunc{
	return func(ctx *gin.Context) {
		jobs, err:=services.GetAllJobsService(db)
		if err!=nil{
			ctx.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
			return 
		}
		ctx.JSON(http.StatusOK,jobs)
	}
}

func GetJobsByUserIdHandler(db *sql.DB)gin.HandlerFunc{
	return  func(ctx *gin.Context) {
		jobs,err:=services.GetJobsByUserIdService(db,ctx.GetInt("userID"))
		if err!=nil{
			ctx.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
			return 
		}
		ctx.JSON(http.StatusOK,jobs)
	}
}

func GetJobByIdHandler(db *sql.DB)gin.HandlerFunc{
	return  func(ctx *gin.Context) {
		id,err:=strconv.Atoi(ctx.Param("id"))
		if err!=nil{
			ctx.JSON(http.StatusBadRequest,gin.H{"error":"⚠️ Invalid job-ID!"})
			return 
		}
		job,err:=services.GetJobByIdService(db,id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
			return 
		}
		ctx.JSON(http.StatusOK,job)
}
}
