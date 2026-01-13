package auth

import (
	"bytes"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/iamskyy666/golang-job-portal-api/internal/models"
)

func PasswordValidationMiddleWare() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		// Read the request body
		bodyBytes,err:=io.ReadAll(ctx.Request.Body)

		if err!=nil{
			ctx.JSON(400, gin.H{"error":"⚠️ ERROR reading request-body!"})
			ctx.Abort()
			return 
		}

		// Create new reader with the bytes
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// Parse the request
		var req models.ChangePasswordRequest
		if err:=ctx.ShouldBindPlain(&req);err!=nil{
			ctx.JSON(400, gin.H{"error":"⚠️ Invalid request-body!"})
			ctx.Abort()
			return 
		}

		// Restore the req. body for the next mw/handler
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		ctx.Next()
	}
}