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


func ForgotPasswordService(db *sql.DB, username string) (string, error) {
	user, err := repository.GetUserByUserName(db, username)
	if err != nil {
		return "", err
	}

	// 1. Generate plain password
	newPassword := utils.GenerateRandomPassword(6)

	// 2. Hash it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// 3. Save hashed password
	user.Password = string(hashedPassword)
	if err := repository.UpdateUserPasswordRepo(db, user); err != nil {
		return "", err
	}

	// 4. RETURN PLAIN PASSWORD (IMPORTANT)
	return newPassword, nil
}


