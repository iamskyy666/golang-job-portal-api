package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, db *sql.DB){
	// AUTH ROUTES
	r.POST("/login",LoginHandler(db))
	r.POST("/register",RegisterHandler(db))

}