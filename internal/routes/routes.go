package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/iamskyy666/golang-job-portal-api/internal/handlers"
)

func InitRoutes(r *gin.Engine, db *sql.DB){
	// AUTH routes
	r.POST("/login",handlers.LoginHandler(db))
	r.POST("/register",handlers.RegisterHandler(db))

	// USER routes
	r.GET("/users/:id",handlers.GetUserHandler(db))
}