package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamskyy666/golang-job-portal-api/internal/models"
	"github.com/iamskyy666/golang-job-portal-api/internal/services"
)

func LoginHandler(db *sql.DB)gin.HandlerFunc{
	return  func(ctx *gin.Context) {
		//var user models.User
	}
}

func RegisterHandler(db *sql.DB)gin.HandlerFunc{
	return  func(ctx *gin.Context) {
		var user models.User
		if err:=ctx.ShouldBindJSON(&user);err!=nil{
			ctx.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
				"status_code":http.StatusBadRequest,
			})
			return 
		}

		err := services.RegisterUser(db,&user)
		if err!=nil{
			ctx.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
				"status_code":http.StatusInternalServerError,
			})
			return 
		}
		ctx.JSON(http.StatusCreated,gin.H{
			"message":"User created/registered successfully! ✅",
			"status_code":http.StatusCreated,
		})
		
	}
}