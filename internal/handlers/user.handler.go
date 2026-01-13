package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
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
				"status_code":http.StatusBadRequest,
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

func UpdateUserProfileHandler(db *sql.DB)gin.HandlerFunc{
	return func(ctx *gin.Context) {
		id,err:=strconv.Atoi(ctx.Param("id"))
		if err!=nil{
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":"Invalid user-ID!",
				"status_code":http.StatusBadRequest,
			})
			return 
		}

		var userUpdate struct{
			Username string `json:"username"`
			Email string `json:"email"`
		}

		if err:=ctx.ShouldBindJSON(&userUpdate);err!=nil{
			ctx.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
				"status_code":http.StatusBadRequest,
			})
			return 
		}

		userID:=ctx.GetInt("userID")
		isAdmin:= ctx.GetBool("isAdmin")

		if !isAdmin && userID!=id{
			ctx.JSON(http.StatusUnauthorized,gin.H{"error":"⚠️ Unauthorized!"})
			return 
		}

		updatedUser,err:=services.UpdateUserProfile(db,id, userUpdate.Username,userUpdate.Email)

		if err!=nil{
			ctx.JSON(http.StatusInternalServerError,gin.H{"error":"⚠️Error updating user-profile"})
			return 
		}

		ctx.JSON(http.StatusOK,updatedUser)
	}
}

func UpdateProfilePictureHandler(db *sql.DB)gin.HandlerFunc{
	return func(ctx *gin.Context) {
		id,err:=strconv.Atoi(ctx.Param("id"))
		if err!=nil{
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":"Invalid user-ID!",
				"status_code":http.StatusBadRequest,
			})
			return 
		}

		userID:=ctx.GetInt("userID")
		isAdmin:= ctx.GetBool("isAdmin")

		if !isAdmin && userID!=id{
			ctx.JSON(http.StatusUnauthorized,gin.H{"error":"⚠️ Unauthorized!"})
			return 
		}

		file,err:=ctx.FormFile("profile_picture")
		if err!=nil{
			ctx.JSON(http.StatusBadRequest,gin.H{"error":"⚠️Error updating profile-picture!"})
			return 
		}

		if err:=os.MkdirAll(os.Getenv("UPLOAD_DIR"),os.ModePerm);err!=nil{
			ctx.JSON(http.StatusUnauthorized,gin.H{"error":"⚠️ ERROR creating upload-directory!"})
			return 
		}

		filename:=fmt.Sprintf("%d-%s",id,filepath.Base(file.Filename))
		filePath:=filepath.Join(os.Getenv("UPLOAD_DIR"),filename)

		if err:=ctx.SaveUploadedFile(file,filePath); err!= nil {
			ctx.JSON(http.StatusInternalServerError,gin.H{"error":"⚠️ ERROR saving uploaded-file!"})
			return 
		}

		err = services.UpdateProfilePicture(db,id,filename)
		if err!=nil{
			ctx.JSON(http.StatusInternalServerError,gin.H{"error":"⚠️ERROR updating profile-picture!"})
			return 
		}

	}
}

func GetUsersHandler(db *sql.DB)gin.HandlerFunc{
	return func(ctx *gin.Context) {
		// check whether user is admin or not
		isAdmin:=ctx.GetBool("isAdmin")
		if(isAdmin==false){
			ctx.JSON(http.StatusUnauthorized,gin.H{"error":"⚠️ Unauthorized to fetch all users!"})
			return 
		}
		users, err:=services.GetUsersService(db)
		if err!=nil{
			ctx.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
			return 
		}
		ctx.JSON(http.StatusOK,users)
	}
}


func DeleteUserHandler(db *sql.DB)gin.HandlerFunc{
	return func(ctx *gin.Context) {
		// check whether user is admin or not
		isAdmin:=ctx.GetBool("isAdmin")
		if(isAdmin==false){
			ctx.JSON(http.StatusUnauthorized,gin.H{"error":"⚠️ Unauthorized to fetch all users!"})
			return 
		}

		id,err:=strconv.Atoi(ctx.Param("id"))
		if err!=nil{
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":"Invalid user-ID!",
				"status_code":http.StatusBadRequest,
			})
			return 
		}

		// check if current user is trynna delete himself/herself
		cuuUserID:=ctx.GetInt("userID")
		if cuuUserID==id{
			ctx.JSON(http.StatusBadRequest,gin.H{"ERROR":"⚠️ You cannot delete yourself!"})
			return 
		}

		// Delete User
		err = services.DeleteUserService(ctx, db,id)

		if err!=nil{
		if err.Error() == "user not found"{
			ctx.JSON(http.StatusNotFound,gin.H{"ERROR":"⚠️ User not found!"})
			return 
		}
		ctx.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return 
	 }
	 ctx.JSON(http.StatusOK,gin.H{"message":"User and associated data deleted successfully! ☑️"})
	}
}