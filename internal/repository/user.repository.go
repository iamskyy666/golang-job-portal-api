package repository

import (
	"database/sql"
	"log"

	"github.com/iamskyy666/golang-job-portal-api/internal/models"
)

func CreateUser(db *sql.DB, user *models.User)error{
	_,err:=db.Exec(`INSERT INTO users (username, password, email) VALUES (?,?,?)`,user.Username,user.Password,user.Email)
	if err != nil {
		return err
	}
	return nil
}

func GetUserById(db *sql.DB, id int)(*models.User,error){
	var user models.User
	var profilePicture sql.NullString // Use sql.NullString to handle NULL values
	err:=db.QueryRow(`SELECT * FROM users WHERE id = ?`,id).Scan(
    &user.ID,
    &user.Username,
    &user.Password,
    &user.Email,
    &user.CreatedAt,
    &user.UpdatedAt,
    &user.IsAdmin,
    &profilePicture,
)
	if err != nil {
		log.Println("⚠️ ERROR:",err.Error())
		return nil, err
	}
	if profilePicture.Valid{
		user.ProfilePicture = &profilePicture.String
	}else{
		user.ProfilePicture = nil
	}
	return &user,nil
}

func GetUserByUserName(db *sql.DB, username string)(*models.User, error){
	user:=&models.User{}

	err:=db.QueryRow(`SELECT * FROM users WHERE username = ?`,username).Scan(
    &user.ID,
    &user.Username,
    &user.Password,
    &user.Email,
    &user.CreatedAt,
    &user.UpdatedAt,
    &user.IsAdmin,
    &user.ProfilePicture,
)

if err != nil {
		log.Println("⚠️ ERROR:",err.Error())
		return nil, err
	}
return user,nil
}