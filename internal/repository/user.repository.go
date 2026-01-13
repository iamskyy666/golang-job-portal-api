package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/iamskyy666/golang-job-portal-api/internal/models"
	"golang.org/x/crypto/bcrypt"
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

func UpdateUser(db *sql.DB, user *models.User)(*models.User, error){
	_,err:=db.Exec("UPDATE users SET username = ?, email = ? WHERE id = ?",user.Username,user.Email, user.ID)
	if err != nil {
		log.Println("ERROR:",err.Error())
		return nil, err
	}

	return user, nil
}

func UpdateProfilePic(db *sql.DB, id int, profilePicture string)error{
	_,err:=db.Exec("UPDATE users SET profile_picture = ? WHERE id = ?",profilePicture, id)
	if err != nil {
		log.Println("ERROR:",err.Error())
		return err
	}
	return nil
}

func GetUsersRepo(db *sql.DB)([]*models.User,error){
rows, err := db.Query("SELECT * FROM users")

	var users []*models.User
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		var profilePicture sql.NullString
		if err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.IsAdmin,
			&profilePicture,
		); err != nil {
			return nil, err
		}

		if profilePicture.Valid{
			user.ProfilePicture = &profilePicture.String
		}else{
			user.ProfilePicture = nil
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}


func UpdateUserPasswordRepo(db *sql.DB,user *models.User)error{
	_,err:=db.Exec("UPDATE users SET password = ? WHERE id = ?",user.Password, user.ID)
	if err != nil {
		log.Println("ERROR:",err)
		return err
	}
	return nil
}

func DeleteUserWithTransactionRepo(tx *sql.Tx, userID int)(string,error){
	// Delete associated jobs first
	_,err:=tx.Exec("DELETE FROM jobs WHERE user_id = ?",userID)
	if err != nil {
		log.Println("ERROR:",err.Error())
		return "",fmt.Errorf("ERROR deleting user's jobs: %v",err)
	}

	// Get user's profile before deleting
	var profilePicture sql.NullString
	err=tx.QueryRow("SELECT profile_picture FROM users WHERE id = ?",userID).Scan(&profilePicture)
	if err != nil {
		log.Println("ERROR:",err.Error())
		return "",fmt.Errorf("ERROR fetching user's profile_picture: %v",err)
	}

	// Delete the user
	result,err:=tx.Exec("DELETE FROM users WHERE id = ?",userID)
	if err != nil {
		log.Println("ERROR:",err.Error())
		return "",fmt.Errorf("ERROR deleting User: %v",err)
	}

	rowsAffected,err:=result.RowsAffected()
	if err != nil {
		log.Println("ERROR:",err.Error())
		return "",fmt.Errorf("ERROR getting rows-affected: %v",err)
	}

	if rowsAffected==0{
		return "",sql.ErrNoRows
	}

	return profilePicture.String,nil
}


func ChangePasswordRepo(db *sql.DB, userID int,currentPassword string, newPassword string)error{
	// First fetch and validate curr. password
	var hashedPassword string
	err:=db.QueryRow("SELECT password FROM users WHERE id = ?",userID).Scan(&hashedPassword)
	if err!=nil{
		if err==sql.ErrNoRows{
			return fmt.Errorf("user_id not found : %v",err)
		}
		return fmt.Errorf("ERROR changing the password : %v",err)
	}

	// Verify whether curr. password is correct or not
	if err:=bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(currentPassword));err!=nil{
		return fmt.Errorf("ERROR! Current pasword is incorrect.")
	}

	// Only if new password is correct, proceed to update
	hashedNewPassword,err:=bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err!=nil{
		return fmt.Errorf("ERROR hashing new password: %v",err)
	}

	result, err:= db.Exec("UPDATE users SET password = ? WHERE id = ?",hashedNewPassword,userID)
	if err!=nil{
		return fmt.Errorf("ERROR updating password: %v",err)
	}

	rowsAffected,err:=result.RowsAffected()
	if err != nil {
		log.Println("ERROR:",err.Error())
		return fmt.Errorf("ERROR checking update result: %v",err)
	}

	if rowsAffected==0{
		return fmt.Errorf("No User found with id: %d",userID)
	}
	return nil
}