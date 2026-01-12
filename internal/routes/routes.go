package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/iamskyy666/golang-job-portal-api/internal/auth"
	"github.com/iamskyy666/golang-job-portal-api/internal/handlers"
)

func InitRoutes(r *gin.Engine, db *sql.DB){
	// AUTH routes
	r.POST("/login",handlers.LoginHandler(db))
	r.POST("/register",handlers.RegisterHandler(db))
	r.GET("/jobs",handlers.GetAllJobsHandler(db))

	// USER routes
	authenticated:=r.Group("/")
	authenticated.Use(auth.AuthMiddleware())
	authenticated.GET("/users/:id",handlers.GetUserHandler(db))
	authenticated.PUT("/users/:id",handlers.UpdateUserProfileHandler(db))
	authenticated.POST("/users/:id",handlers.UpdateProfilePictureHandler(db))

	//older way: r.PUT("/users/:id",auth.AuthMiddleware(),handlers.UpdateUserProfileHandler(db))

	// JOB routes
	authenticated.POST("/jobs",handlers.CreateJobHandler(db))
	authenticated.GET("/jobs-by-user",handlers.GetJobsByUserIdHandler(db))
	authenticated.GET("/jobs/:id",handlers.GetJobByIdHandler(db))
	authenticated.PUT("/jobs/:id",handlers.UpdateJobHandler(db))
}