package services

import (
	"database/sql"
	"log"

	"github.com/iamskyy666/golang-job-portal-api/internal/models"
	"github.com/iamskyy666/golang-job-portal-api/internal/repository"
	"github.com/iamskyy666/golang-job-portal-api/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(db *sql.DB, user *models.User)error{
	hashedPassword,err:= bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return repository.CreateUser(db,user)
}

func LoginUser(db *sql.DB, username, password string)(string, error){
	user,err:=repository.GetUserByUserName(db,username)
	if err!=nil{
		log.Println("ERROR:",err.Error())
		return "",err
	}

	// check if password is matching
	if err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password));err!=nil{
		log.Println("ERROR:",err.Error())
		return "",err
	}
	
	return utils.GenerateToken(user.Username, user.ID, user.IsAdmin)
}



