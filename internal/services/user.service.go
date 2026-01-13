package services

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/iamskyy666/golang-job-portal-api/internal/models"
	"github.com/iamskyy666/golang-job-portal-api/internal/repository"
	"github.com/iamskyy666/golang-job-portal-api/pkg/utils"
)

func GetUserById(db *sql.DB, id int) (*models.User,error) {
	return repository.GetUserById(db,id)
}

func UpdateUserProfile(db *sql.DB, id int, username, emailId string)(*models.User,error){
	user:=&models.User{ID:id,Username: username, Email: emailId}
	return repository.UpdateUser(db,user)
}

func UpdateProfilePicture(db *sql.DB,id int, profilePicture string)error{
	return repository.UpdateProfilePic(db,id,profilePicture)
}

func GetUsersService(db *sql.DB)([]*models.User,error){
	return repository.GetUsersRepo(db)
}

func DeleteUserService(ctx *gin.Context, db *sql.DB, userID int)error{
	// start a transaction
	tx,err:=db.BeginTx(ctx.Request.Context(),&sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return fmt.Errorf("Error starting transaction: %v",err)
	}

	defer tx.Rollback() // rollback, if not committed

	// Delete user and associated data from 'users' table
	profilePicture,err:=repository.DeleteUserWithTransactionRepo(tx,userID)
	if err != nil {
		if err==sql.ErrNoRows{
			return fmt.Errorf("ERROR - User not found: %v",err)
		}
		return fmt.Errorf("Error deleting User: %v",err)
	}

	// Commit the tx
	if err=tx.Commit();err!=nil{
		return fmt.Errorf("Error committing transaction: %v",err)
	}

	// Delete profile picture after successful transaction if it exists
	if profilePicture!=""{
		filePath:=filepath.Join(os.Getenv("UPLOAD_DIR"),profilePicture)
		err = utils.DeleteFileIfExists(filePath) // f(x) in utils file.
		if err!=nil{
			return fmt.Errorf("ERROR deleting profile picture: %v",err)
		}
	}

	return nil

}