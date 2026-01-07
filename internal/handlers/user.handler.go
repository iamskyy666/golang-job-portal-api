package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iamskyy666/golang-job-portal-api/internal/services"
)

func GetUserHandler(db *sql.DB)gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id,err:=strconv.Atoi(ctx.Param("id"))
		if err!=nil{
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":"Invalid user-ID",
				"status_code":http.StatusBadGateway,
			})
			return 
		}
		user,err:=services.GetUserById(db,id)

		if err!=nil{
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":err.Error(),
				"status_code":http.StatusInternalServerError,
			})
			return 
		}

		ctx.JSON(http.StatusOK, user)
	}
}