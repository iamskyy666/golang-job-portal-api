package handlers

import (
	"database/sql"
	"net/http"

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