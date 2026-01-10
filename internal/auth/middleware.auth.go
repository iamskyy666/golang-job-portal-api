package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/iamskyy666/golang-job-portal-api/pkg/utils"
)

func AuthMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		token:=ctx.GetHeader("Authorization")
		if token==""{
			ctx.AbortWithStatusJSON(401,gin.H{"error":"⚠️ Missing Authorization header!"})
			ctx.Abort()
			return
		}
		claims,err:=utils.ValidateToken(token)

		if err!=nil{
			ctx.AbortWithStatusJSON(401,gin.H{"error":"⚠️ Invalid token!"})
			ctx.Abort()
			return
		}

		ctx.Set("userID",claims.UserID)
		ctx.Set("isAdmin",claims.IsAdmin)

		ctx.Next()
	}
}